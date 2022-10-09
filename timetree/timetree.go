package timetree

import (
	"net/http"
	"net/url"
	"sync"
)

const (
	defaultBaseURL = "https://timetreeapis.com"
)

type Client struct {
	clientMu sync.Mutex
	client   *http.Client
	BaseURL  *url.URL
	Token    string
}

func NewClient(httpclient *http.Client, token string) *Client {
	if httpclient == nil {
		httpclient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:  httpclient,
		BaseURL: baseURL,
		Token:   token,
	}
	return c
}
