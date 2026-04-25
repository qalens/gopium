package gopium

import (
	"context"
	"net/http"
)

func (d *Driver) ExecuteScript(ctx context.Context, script string, args ...any) (any, error) {
	envelope, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/execute/sync"), map[string]any{
		"script": script,
		"args":   args,
	})
	if err != nil {
		return nil, err
	}
	return envelope.Value, nil
}

func (d *Driver) ExecuteAsyncScript(ctx context.Context, script string, args ...any) (any, error) {
	envelope, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/execute/async"), map[string]any{
		"script": script,
		"args":   args,
	})
	if err != nil {
		return nil, err
	}
	return envelope.Value, nil
}

func (d *Driver) ExecuteMobile(ctx context.Context, method string, args map[string]any) (any, error) {
	if args == nil {
		args = map[string]any{}
	}
	return d.ExecuteScript(ctx, "mobile: "+method, args)
}
