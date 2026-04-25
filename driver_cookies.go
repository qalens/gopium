package gopium

import (
	"context"
	"fmt"
	"net/http"
)

func (d *Driver) Cookies(ctx context.Context) ([]Cookie, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath("/cookie"), nil)
	if err != nil {
		return nil, err
	}
	var cookies []Cookie
	if err := remarshal(envelope.Value, &cookies); err != nil {
		return nil, fmt.Errorf("decode cookies: %w", err)
	}
	return cookies, nil
}

func (d *Driver) Cookie(ctx context.Context, name string) (Cookie, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath("/cookie/"+name), nil)
	if err != nil {
		return Cookie{}, err
	}
	var cookie Cookie
	if err := remarshal(envelope.Value, &cookie); err != nil {
		return Cookie{}, fmt.Errorf("decode cookie: %w", err)
	}
	return cookie, nil
}

func (d *Driver) AddCookie(ctx context.Context, cookie Cookie) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/cookie"), map[string]any{"cookie": cookie})
	return err
}

func (d *Driver) DeleteCookie(ctx context.Context, name string) error {
	_, err := d.client.do(ctx, http.MethodDelete, d.sessionPath("/cookie/"+name), nil)
	return err
}

func (d *Driver) DeleteAllCookies(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodDelete, d.sessionPath("/cookie"), nil)
	return err
}
