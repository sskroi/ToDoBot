package telegram

import (
	"ToDoBot1/pkg/e"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

// New возвращает объект для взаимодействия с API telegram
func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
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
	err := c.sendMessage(chatId, text, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendMessageReplyMarup(chatId uint64, text string, replyMarkup *ReplyKeyboardMarkup) error {
	err := c.sendMessage(chatId, text, replyMarkup)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) sendMessage(chatId uint64, text string, replyMarkup *ReplyKeyboardMarkup) error {
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

	// добавляем к объекту request параметры запроса
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
