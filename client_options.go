package gopium

import (
	"net/http"
	"strings"
)

func WithHTTPClient(httpClient HTTPDoer) ClientOption {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		if strings.TrimSpace(userAgent) != "" {
			c.userAgent = userAgent
		}
	}
}

func WithRequestEditor(editor RequestEditor) ClientOption {
	return func(c *Client) {
		if editor != nil {
			c.editors = append(c.editors, editor)
		}
	}
}

func WithBasicAuth(username, password string) ClientOption {
	return WithRequestEditor(func(req *http.Request) {
		req.SetBasicAuth(username, password)
	})
}

func WithLogger(logger Logger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}
