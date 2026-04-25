package gopium

import (
	"context"
	"net/http"
)

func (d *Driver) PerformActions(ctx context.Context, actions []Action) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/actions"), map[string]any{"actions": actions})
	return err
}

func (d *Driver) ReleaseActions(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodDelete, d.sessionPath("/actions"), nil)
	return err
}
