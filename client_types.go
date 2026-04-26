package gopium

import (
	"net/http"
	"net/url"
)

type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type RequestEditor func(*http.Request)

type ClientOption func(*Client)

type Logger interface {
	Printf(format string, v ...any)
}

type Client struct {
	baseURL    *url.URL
	httpClient HTTPDoer
	userAgent  string
	editors    []RequestEditor
	logger     Logger
}
