package gopium

import (
	"context"
	"net/http"
	"strings"
)

func (d *Driver) SessionCommand(ctx context.Context, method, suffix string, payload any) (any, error) {
	if !strings.HasPrefix(suffix, "/") {
		suffix = "/" + suffix
	}
	envelope, err := d.client.do(ctx, method, d.sessionPath(suffix), payload)
	if err != nil {
		return nil, err
	}
	return envelope.Value, nil
}

func (d *Driver) AppiumCommand(ctx context.Context, suffix string, payload any) (any, error) {
	return d.SessionCommand(ctx, http.MethodPost, "/appium/"+strings.TrimPrefix(suffix, "/"), payload)
}

func (d *Driver) ExecuteCDP(ctx context.Context, vendor, cmd string, params map[string]any) (any, error) {
	if params == nil {
		params = map[string]any{}
	}
	return d.SessionCommand(ctx, http.MethodPost, "/"+vendor+"/cdp/execute", map[string]any{
		"cmd":    cmd,
		"params": params,
	})
}
