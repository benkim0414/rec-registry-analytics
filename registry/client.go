package registry

import (
	"github.com/valyala/fasthttp"
)

// Client is a HTTP client that is provided by fasthttp package.
type Client struct {
	hc *fasthttp.Client
}

// NewClient returns a *Client which wraps *fasthttp.Client.
func NewClient() *Client {
	return &Client{
		hc: &fasthttp.Client{},
	}
}

// Do performs the given HTTP request and fills the HTTP response.
func (c *Client) Do(req *fasthttp.Request) (resp *fasthttp.Response, err error) {
	resp = fasthttp.AcquireResponse()
	err = c.hc.Do(req, resp)
	return
}
