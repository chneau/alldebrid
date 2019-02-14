package alldebrid

import (
	"encoding/json"
	"log"
	"net/http"
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

// Connect connects to get the token
func (c *Client) Connect(username, password string) (*LoginResponse, error) {
	res, err := c.HTTPClient.Do(c.buildReq("user/login", k{"username": username, "password": password}))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	loginResponse := &LoginResponse{}
	err = json.NewDecoder(res.Body).Decode(loginResponse)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c.Token = loginResponse.Token
	return loginResponse, nil
}

// GetDownloadLink ...
func (c *Client) GetDownloadLink(link string) (*LinkUnlockResponse, error) {
	res, err := c.HTTPClient.Do(c.buildReq("link/unlock", k{"link": link}))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	linkUnlockResponse := &LinkUnlockResponse{}
	err = json.NewDecoder(res.Body).Decode(linkUnlockResponse)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return linkUnlockResponse, nil
}

// New returns an instance of the client.
func New() *Client {
	return &Client{
		Base:       "https://api.alldebrid.com/",
		Agent:      "A Go client",
		HTTPClient: http.DefaultClient,
	}
}
