package gopium

import (
	"context"
	"fmt"
	"net/http"
)

func (d *Driver) SetTimeouts(ctx context.Context, timeouts Timeouts) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/timeouts"), timeouts)
	return err
}

func (d *Driver) GetTimeouts(ctx context.Context) (Timeouts, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath("/timeouts"), nil)
	if err != nil {
		return Timeouts{}, err
	}

	var result Timeouts
	if err := remarshal(envelope.Value, &result); err != nil {
		return Timeouts{}, fmt.Errorf("decode timeouts: %w", err)
	}
	return result, nil
}
