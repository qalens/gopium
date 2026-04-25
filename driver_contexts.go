package gopium

import (
	"context"
	"fmt"
	"net/http"
)

func (d *Driver) CurrentContext(ctx context.Context) (string, error) {
	return d.readString(ctx, http.MethodGet, d.sessionPath("/context"), nil)
}

func (d *Driver) SetContext(ctx context.Context, name string) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/context"), map[string]any{"name": name})
	return err
}

func (d *Driver) Contexts(ctx context.Context) ([]string, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath("/contexts"), nil)
	if err != nil {
		return nil, err
	}
	items, ok := envelope.Value.([]any)
	if !ok {
		return nil, fmt.Errorf("expected context list but got %T", envelope.Value)
	}
	out := make([]string, 0, len(items))
	for _, item := range items {
		text, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("expected context name string but got %T", item)
		}
		out = append(out, text)
	}
	return out, nil
}
