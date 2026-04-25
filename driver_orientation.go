package gopium

import (
	"context"
	"fmt"
	"net/http"
)

func (d *Driver) Orientation(ctx context.Context) (string, error) {
	return d.readString(ctx, http.MethodGet, d.sessionPath("/orientation"), nil)
}

func (d *Driver) SetOrientation(ctx context.Context, orientation string) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/orientation"), map[string]any{"orientation": orientation})
	return err
}

func (d *Driver) Rotation(ctx context.Context) (Rotation, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath("/rotation"), nil)
	if err != nil {
		return Rotation{}, err
	}
	var rotation Rotation
	if err := remarshal(envelope.Value, &rotation); err != nil {
		return Rotation{}, fmt.Errorf("decode rotation: %w", err)
	}
	return rotation, nil
}

func (d *Driver) SetRotation(ctx context.Context, rotation Rotation) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/rotation"), rotation)
	return err
}
