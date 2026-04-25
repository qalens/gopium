package gopium

import (
	"context"
	"net/http"
)

func (d *Driver) StartRecordingScreen(ctx context.Context, options RecordScreenOptions) error {
	payload := map[string]any{}
	if len(options) > 0 {
		payload["options"] = map[string]any(options)
	}
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/start_recording_screen"), payload)
	return err
}

func (d *Driver) StopRecordingScreen(ctx context.Context, options RecordScreenOptions) (string, error) {
	payload := map[string]any{}
	if len(options) > 0 {
		payload["options"] = map[string]any(options)
	}
	return d.postString(ctx, "/appium/stop_recording_screen", payload)
}
