package gopium

import (
	"net"
	"net/http"
	"strings"
	"time"
)

func WithHTTPClient(httpClient HTTPDoer) ClientOption {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

func WithHTTPTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		if timeout <= 0 {
			return
		}
		if client, ok := c.httpClient.(*http.Client); ok && client != nil {
			client.Timeout = timeout
			if transport, ok := client.Transport.(*http.Transport); ok && transport != nil {
				transport.TLSHandshakeTimeout = timeout
				transport.ResponseHeaderTimeout = timeout
				if dialContext := transport.DialContext; dialContext != nil {
					transport.DialContext = (&net.Dialer{
						Timeout:   timeout,
						KeepAlive: 30 * time.Second,
					}).DialContext
				}
			}
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
