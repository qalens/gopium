package gopium

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewClient(rawURL string, opts ...ClientOption) (*Client, error) {
	if strings.TrimSpace(rawURL) == "" {
		rawURL = defaultServerURL
	}

	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("parse server url: %w", err)
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return nil, fmt.Errorf("invalid server url %q", rawURL)
	}

	client := &Client{
		baseURL: parsed,
		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   60 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   60 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				ResponseHeaderTimeout: 60 * time.Second,
			},
			Timeout: 60 * time.Second,
		},
		userAgent: "gopium/0.1",
	}

	for _, opt := range opts {
		if opt != nil {
			opt(client)
		}
	}

	return client, nil
}

func MustNewClient(rawURL string, opts ...ClientOption) *Client {
	client, err := NewClient(rawURL, opts...)
	if err != nil {
		panic(err)
	}
	return client
}

func (c *Client) NewDriver(ctx context.Context, options SessionOptions) (*Driver, error) {
	driver := &Driver{client: c}
	if err := driver.StartSession(ctx, options); err != nil {
		return nil, err
	}
	return driver, nil
}

func (c *Client) Status(ctx context.Context) (map[string]any, error) {
	envelope, err := c.do(ctx, http.MethodGet, "/status", nil)
	if err != nil {
		return nil, err
	}
	return envelope.valueAsMap()
}

func (c *Client) resolve(path string) string {
	base := strings.TrimRight(c.baseURL.String(), "/")
	switch {
	case path == "":
		return base
	case strings.HasPrefix(path, "/"):
		return base + path
	default:
		return base + "/" + path
	}
}
