package telegram

import (
	"ToDoBot1/pkg/e"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
    deleteWebhookMethod = "deleteWebhook"
    SetWebhookMethod = "setWebhook"
)

type Config struct {
	Token string `toml:"token"`
	Host  string `toml:"host"`
}

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(cfg Config) *Client {
	return &Client{
		host:     cfg.Host,
		basePath: "bot" + cfg.Token,
		client:   http.Client{},
	}
}

func (c *Client) DeleteWebhook() error {
    _, err := c.doRequest(deleteWebhookMethod, url.Values{})
   if err != nil {
       return e.Wrap("can't delete webhook: ", err)
   }

   return nil
}

func (c *Client) SetWebhook(hookURL, certPath string) error {
    file, err := os.Open(certPath)
    if err != nil {
        return e.Wrap("can't set webhook: ", err)
    }
    defer file.Close()
    
    fileContent, err := io.ReadAll(file)
    if err != nil {
        return err
    }

    fileInfo, err := file.Stat()
    if err != nil {
        return err
    }

    body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("certificate", fileInfo.Name())
    if err != nil {
        return err
    }
    if _, err := part.Write(fileContent); err != nil {
        return err
    }

    if err := writer.WriteField("url", hookURL); err != nil {
        return err
    }
    
    url := url.URL{
        Scheme: "https",
        Host: c.host,
        Path: path.Join(c.basePath, SetWebhookMethod),
    }

    t := writer.FormDataContentType()

    if err := writer.Close(); err != nil {
        return err
    }


    resp, err := http.Post(url.String(), t, body)
    if err != nil {
        return e.Wrap("can't set webhook: ", err)
    }

    fmt.Println(url.String(), writer.FormDataContentType(), body.String()) // debug info


    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    log.Println("telegram setWebhook answer: ", resp.StatusCode, string(respBody))

    return nil
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	const errMsg = "can't get updates"

	querryParams := url.Values{}
	querryParams.Add("offset", strconv.Itoa(offset))
	querryParams.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, querryParams)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	var res UpdatesResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	return res.Result, nil
}

func (c *Client) SendMessage(chatId uint64, text string) error {
	err := c.sendMsg(chatId, text, nil)
	if err != nil {
		return err
	}

	return nil
}

// SendMessageRM sens message with `reply_markup` parameter
func (c *Client) SendMessageRM(chatId uint64, text string, replyMarkup interface{}) error {
	err := c.sendMsg(chatId, text, replyMarkup)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) sendMsg(chatId uint64, text string, replyMarkup interface{}) error {
	querryParams := url.Values{}
	querryParams.Add("chat_id", strconv.FormatUint(chatId, 10))
	querryParams.Add("text", text)

	querryParams.Add("parse_mode", "HTML")

	if replyMarkup != nil {
		serializedReplyMarkup, err := json.Marshal(replyMarkup)
		if err != nil {
			return e.Wrap("can't send message", err)
		}

		querryParams.Add("reply_markup", string(serializedReplyMarkup))
	}

	_, err := c.doRequest(sendMessageMethod, querryParams)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil
}

func (c *Client) doRequest(method string, querryParams url.Values) ([]byte, error) {
	const errMsg = "can't do request"

	url := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	request.URL.RawQuery = querryParams.Encode()

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	return body, nil
}
