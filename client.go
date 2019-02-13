package alldebrid

import (
	"encoding/json"
	"log"
	"net/http"
)

// Client ...
type Client struct {
	Base       string
	Agent      string
	HTTPClient *http.Client
	Token      string
}

func (c *Client) buildReq(path string, queries map[string]string) *http.Request {
	req, _ := http.NewRequest("GET", c.Base+"user/login", nil)
	q := req.URL.Query()
	q.Set("agent", c.Agent)
	for k, v := range queries {
		q.Set(k, v)
	}
	if c.Token == "" {
		q.Set("token", c.Token)
	}
	req.URL.RawQuery = q.Encode()
	return req
}

// Connect connects to get the token
func (c *Client) Connect(username, password string) (*LoginResponse, error) {
	req := c.buildReq("user/login", map[string]string{"username": username, "password": password})
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	loginResponse := LoginResponse{}
	err = json.NewDecoder(res.Body).Decode(&loginResponse)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c.Token = loginResponse.Token
	return &loginResponse, nil
}

// New returns an instance of the client.
func New() *Client {
	return &Client{
		Base:       "https://api.alldebrid.com/",
		Agent:      "A Go client",
		HTTPClient: http.DefaultClient,
	}
}
