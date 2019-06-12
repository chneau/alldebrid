package alldebrid

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type k map[string]string // aliasing

// Client ...
type Client struct {
	Base       string
	Agent      string
	HTTPClient *http.Client
	Token      string
}

func (c *Client) buildReq(path string, queries k) *http.Request {
	req, _ := http.NewRequest("GET", c.Base+path, nil)
	q := req.URL.Query()
	q.Set("agent", c.Agent)
	for k, v := range queries {
		q.Add(k, v)
	}
	q.Set("token", c.Token)
	req.URL.RawQuery = q.Encode()
	return req
}

// GetDownloadLink ...
func (c *Client) GetDownloadLink(link string) (*LinkUnlock, error) {
	res, err := c.HTTPClient.Do(c.buildReq("link/unlock", k{"link": link}))
	if err != nil {
		return nil, err
	}
	linkUnlock := &LinkUnlock{}
	err = json.NewDecoder(res.Body).Decode(linkUnlock)
	if err != nil {
		return nil, err
	}
	if linkUnlock.Error != nil {
		return nil, errors.New(*linkUnlock.Error)
	}
	return linkUnlock, nil
}

// GetPin ...
func (c *Client) GetPin() (*Pin, error) {
	res, err := c.HTTPClient.Do(c.buildReq("pin/get", k{}))
	if err != nil {
		return nil, err
	}
	pin := &Pin{}
	err = json.NewDecoder(res.Body).Decode(pin)
	if err != nil {
		return nil, err
	}
	if pin.Error != nil {
		return nil, errors.New(*pin.Error)
	}
	return pin, nil
}

// CheckPin ...
func (c *Client) CheckPin(pin *Pin) error {
	res, err := c.HTTPClient.Get(pin.CheckURL)
	if err != nil {
		return err
	}
	token := &struct {
		Success   bool    `json:"success,omitempty"`
		Token     string  `json:"token,omitempty"`
		Activated bool    `json:"activated,omitempty"`
		ExpiresIn int64   `json:"expires_in,omitempty"`
		Error     *string `json:"error,omitempty"`
	}{}
	err = json.NewDecoder(res.Body).Decode(token)
	if err != nil {
		return err
	}
	if token.Error != nil {
		return errors.New(*token.Error)
	}
	c.Token = token.Token
	return nil
}

// New returns an instance of the client.
func New() *Client {
	return &Client{
		Base:       "https://api.alldebrid.com/",
		Agent:      "GoClient",
		HTTPClient: http.DefaultClient,
	}
}

// LinkUnlock ...
type LinkUnlock struct {
	Success bool `json:"success,omitempty"`
	Infos   struct {
		Link      string      `json:"link,omitempty"`
		Host      string      `json:"host,omitempty"`
		Filename  string      `json:"filename,omitempty"`
		Streaming interface{} `json:"streaming,omitempty"`
		Paws      bool        `json:"paws,omitempty"`
	} `json:"infos,omitempty"`
	Error *string `json:"error,omitempty"`
}

// Pin ...
type Pin struct {
	Success   bool    `json:"success,omitempty"`
	Pin       string  `json:"pin,omitempty"`
	ExpiredIn int64   `json:"expired_in,omitempty"`
	UserURL   string  `json:"user_url,omitempty"`
	BaseURL   string  `json:"base_url,omitempty"`
	CheckURL  string  `json:"check_url,omitempty"`
	Error     *string `json:"error,omitempty"`
}
