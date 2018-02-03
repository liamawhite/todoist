package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	protocol          = "https://"
	todoistHost       = "beta.todoist.com/"
	todoistPathPrefix = "API/v8/"
)

type Client struct {
	http  *http.Client
	token string
}

func NewClient(token string) *Client {
	return &Client{
		http:  &http.Client{},
		token: token,
	}
}

func (c *Client) do(req *http.Request) (*http.Response, error) {

	// Create new url with host, protocol and API version
	req.URL, _ = url.Parse(protocol + todoistHost + todoistPathPrefix + req.URL.Path)

	// Add token
	queryStr := req.URL.Query()
	queryStr.Add("token", c.token)
	req.URL.RawQuery = queryStr.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if len(body) > 0 {
			return nil, fmt.Errorf("received %v status code: %v", resp.StatusCode, string(body))
		}
	}
	return resp, nil
}
