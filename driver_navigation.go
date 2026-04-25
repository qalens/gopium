package gopium

import (
	"context"
	"net/http"
	"net/url"
)

func (d *Driver) SessionURL() string {
	if d.sessionID == "" {
		return ""
	}

	u, err := url.Parse(d.client.resolve(d.sessionPath("")))
	if err != nil {
		return ""
	}
	return u.String()
}

func (d *Driver) Navigate(ctx context.Context, destination string) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/url"), map[string]any{"url": destination})
	return err
}

func (d *Driver) Back(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/back"), map[string]any{})
	return err
}

func (d *Driver) Forward(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/forward"), map[string]any{})
	return err
}

func (d *Driver) Refresh(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/refresh"), map[string]any{})
	return err
}

func (d *Driver) URL(ctx context.Context) (string, error) {
	return d.readString(ctx, http.MethodGet, d.sessionPath("/url"), nil)
}

func (d *Driver) Title(ctx context.Context) (string, error) {
	return d.readString(ctx, http.MethodGet, d.sessionPath("/title"), nil)
}

func (d *Driver) PageSource(ctx context.Context) (string, error) {
	return d.readString(ctx, http.MethodGet, d.sessionPath("/source"), nil)
}

func (d *Driver) Screenshot(ctx context.Context) (string, error) {
	return d.readString(ctx, http.MethodGet, d.sessionPath("/screenshot"), nil)
}
