package threesixty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zoftko/gowhat/message"
	"io"
	"net/http"
	"net/url"
	"time"
)

var ProductionURL = "https://waba-v2.360dialog.io"
var SandboxURL = "https://waba-sandbox.360dialog.io/v1"

type Client struct {
	client  *http.Client
	token   string
	baseURL string
}

type MessageEntry struct {
	ID string `json:"id"`
}

type Response struct {
	Messages []MessageEntry `json:"messages"`
}

func NewClient(token string, baseURL string) (*Client, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	_, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:  client,
		token:   token,
		baseURL: baseURL,
	}, nil
}

func (c *Client) url(path ...string) string {
	u, _ := url.JoinPath(c.baseURL, path...)
	return u
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("D360-API-Key", c.token)
	req.Header.Set("Content-Type", "application/json")

	return c.client.Do(req)
}

func (c *Client) SendMessage(envelope message.Envelope) (*Response, error) {
	payload, err := json.Marshal(envelope)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url("messages"), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)

		return nil, fmt.Errorf("received status code %d, message: %s", res.StatusCode, body)
	}

	var data Response
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	return &data, nil
}
