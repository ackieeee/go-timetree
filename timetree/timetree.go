package timetree

import (
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
)

const (
	defaultBaseURL = "https://timetreeapis.com"
)

type QueryParam struct {
	*url.Values
}

type FormParam struct {
	*url.Values
}

type Client struct {
	clientMu sync.Mutex
	client   *http.Client
	BaseURL  *url.URL
	Token    string
	Calendar *Calendar
}

type getClient func(spath string, query *QueryParam) (*http.Response, error)
type postClient func(spath string, body *FormParam) (*http.Response, error)

type httpMethod struct {
	Get  getClient
	Post postClient
}

func (c *Client) setRequest(method, spath string, header http.Header, body io.Reader, param *QueryParam) (*http.Request, error) {
	u := c.BaseURL
	u.Path = path.Join(spath)
	if param != nil {
		u.RawQuery = param.Encode()
	}
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	return req, nil
}

func (c *Client) do(method, spath string, header http.Header, body io.Reader, param *QueryParam) (*http.Response, error) {
	req, err := c.setRequest(method, spath, header, body, param)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.timetree.v1+json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	return c.client.Do(req)
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

	m := httpMethod{
		Get: func(spath string, query *QueryParam) (*http.Response, error) {
			return c.do(http.MethodGet, spath, nil, nil, query)
		},
		Post: func(spath string, query *FormParam) (*http.Response, error) {
			return c.do(http.MethodGet, spath, nil, strings.NewReader(query.Encode()), nil)
		},
	}
	c.Calendar = &Calendar{
		httpMethod: m,
	}
	return c
}
