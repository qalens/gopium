package gopium

import (
	"context"
	"fmt"
	"net/http"
)

func (d *Driver) readString(ctx context.Context, method, path string, body any) (string, error) {
	envelope, err := d.client.do(ctx, method, path, body)
	if err != nil {
		return "", err
	}

	value, ok := envelope.Value.(string)
	if !ok {
		return "", fmt.Errorf("expected string response but got %T", envelope.Value)
	}

	return value, nil
}

func (d *Driver) postString(ctx context.Context, path string, body any) (string, error) {
	return d.readString(ctx, http.MethodPost, d.sessionPath(path), body)
}

func (d *Driver) getBool(ctx context.Context, path string) (bool, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath(path), nil)
	if err != nil {
		return false, err
	}
	value, ok := envelope.Value.(bool)
	if !ok {
		return false, fmt.Errorf("expected bool response but got %T", envelope.Value)
	}
	return value, nil
}

func (d *Driver) postBool(ctx context.Context, path string, body any) (bool, error) {
	envelope, err := d.client.do(ctx, http.MethodPost, d.sessionPath(path), body)
	if err != nil {
		return false, err
	}
	value, ok := envelope.Value.(bool)
	if !ok {
		return false, fmt.Errorf("expected bool response but got %T", envelope.Value)
	}
	return value, nil
}
