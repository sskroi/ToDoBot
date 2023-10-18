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
	getUpdatesMethod = "getUpdates"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

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
