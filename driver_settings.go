package gopium

import (
	"context"
	"net/http"
)

func (d *Driver) SetSetting(ctx context.Context, key string, value any) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/settings"), map[string]any{
		"settings": map[string]any{key: value},
	})
	return err
}

func (d *Driver) UpdateSettings(ctx context.Context, settings map[string]any) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/settings"), map[string]any{
		"settings": settings,
	})
	return err
}

func (d *Driver) Settings(ctx context.Context) (map[string]any, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath("/appium/settings"), nil)
	if err != nil {
		return nil, err
	}
	return envelope.valueAsMap()
}
