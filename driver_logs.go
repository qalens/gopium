package gopium

import (
	"context"
	"net/http"
)

func (d *Driver) LogEvent(ctx context.Context, vendor, event string) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/log_event"), map[string]any{
		"vendor": vendor,
		"event":  event,
	})
	return err
}

func (d *Driver) LogEvents(ctx context.Context, eventType string) (map[string]any, error) {
	payload := map[string]any{}
	if eventType != "" {
		payload["type"] = eventType
	}
	envelope, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/events"), payload)
	if err != nil {
		return nil, err
	}
	return envelope.valueAsMap()
}
