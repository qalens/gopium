package gopium

import (
	"context"
	"net/http"
)

func (d *Driver) AcceptAlert(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/alert/accept"), map[string]any{})
	return err
}

func (d *Driver) DismissAlert(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/alert/dismiss"), map[string]any{})
	return err
}

func (d *Driver) AlertText(ctx context.Context) (string, error) {
	return d.readString(ctx, http.MethodGet, d.sessionPath("/alert/text"), nil)
}

func (d *Driver) SetAlertText(ctx context.Context, text string) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/alert/text"), map[string]any{"text": text, "value": splitRunes(text)})
	return err
}
