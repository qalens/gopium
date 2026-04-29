package gopium

import (
	"context"
	"encoding/base64"
	"net/http"
)

func (d *Driver) GetDeviceTime(ctx context.Context, format string) (string, error) {
	payload := map[string]any{}
	if format != "" {
		payload["format"] = format
	}
	return d.postString(ctx, "/appium/device/system_time", payload)
}

func (d *Driver) GetClipboard(ctx context.Context, contentType ClipboardContentType) (string, error) {
	payload := map[string]any{}
	if contentType != "" {
		payload["contentType"] = string(contentType)
	}
	return d.postString(ctx, "/appium/device/get_clipboard", payload)
}

func (d *Driver) GetClipboardText(ctx context.Context) (string, error) {
	encoded, err := d.GetClipboard(ctx, ClipboardPlaintext)
	if err != nil {
		return "", err
	}
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (d *Driver) SetClipboard(ctx context.Context, contentType ClipboardContentType, label, contentBase64 string) error {
	payload := map[string]any{"content": contentBase64}
	if contentType != "" {
		payload["contentType"] = string(contentType)
	}
	if label != "" {
		payload["label"] = label
	}
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/set_clipboard"), payload)
	return err
}

func (d *Driver) SetClipboardText(ctx context.Context, label, text string) error {
	return d.SetClipboard(ctx, ClipboardPlaintext, label, base64.StdEncoding.EncodeToString([]byte(text)))
}

func (d *Driver) PullFile(ctx context.Context, path string) (string, error) {
	return d.postString(ctx, "/appium/device/pull_file", map[string]any{"path": path})
}

func (d *Driver) PullFolder(ctx context.Context, path string) (string, error) {
	return d.postString(ctx, "/appium/device/pull_folder", map[string]any{"path": path})
}

func (d *Driver) PushFile(ctx context.Context, path, dataBase64 string) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/push_file"), map[string]any{
		"path": path,
		"data": dataBase64,
	})
	return err
}

func (d *Driver) PushFileBytes(ctx context.Context, path string, data []byte) error {
	return d.PushFile(ctx, path, base64.StdEncoding.EncodeToString(data))
}

func (d *Driver) HideKeyboard(ctx context.Context, strategy, key string, keyCode int, keyName string) error {
	var options map[string]any
	if strategy != "" {
		if options == nil {
			options = map[string]any{}
		}
		options["strategy"] = strategy
	}
	if key != "" {
		if options == nil {
			options = map[string]any{}
		}
		options["key"] = key
	}
	if keyCode != 0 {
		if options == nil {
			options = map[string]any{}
		}
		options["keyCode"] = keyCode
	}
	if keyName != "" {
		if options == nil {
			options = map[string]any{}
		}
		options["keyName"] = keyName
	}

	var payload any
	if options != nil {
		payload = options
	}

	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/hide_keyboard"), payload)
	return err
}

func (d *Driver) IsKeyboardShown(ctx context.Context) (bool, error) {
	return d.getBool(ctx, "/appium/device/is_keyboard_shown")
}

func (d *Driver) CurrentActivity(ctx context.Context) (string, error) {
	return d.readString(ctx, http.MethodGet, d.sessionPath("/appium/device/current_activity"), nil)
}

func (d *Driver) CurrentPackage(ctx context.Context) (string, error) {
	return d.readString(ctx, http.MethodGet, d.sessionPath("/appium/device/current_package"), nil)
}

func (d *Driver) OpenNotifications(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/open_notifications"), map[string]any{})
	return err
}

func (d *Driver) PressKeyCode(ctx context.Context, keycode int, metastate, flags *int) error {
	payload := map[string]any{"keycode": keycode}
	if metastate != nil {
		payload["metastate"] = *metastate
	}
	if flags != nil {
		payload["flags"] = *flags
	}
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/press_keycode"), payload)
	return err
}

func (d *Driver) LongPressKeyCode(ctx context.Context, keycode int, metastate, flags *int) error {
	payload := map[string]any{"keycode": keycode}
	if metastate != nil {
		payload["metastate"] = *metastate
	}
	if flags != nil {
		payload["flags"] = *flags
	}
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/long_press_keycode"), payload)
	return err
}

func (d *Driver) GetDisplayDensity(ctx context.Context) (int, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath("/appium/device/display_density"), nil)
	if err != nil {
		return 0, err
	}
	switch v := envelope.Value.(type) {
	case float64:
		return int(v), nil
	case int:
		return v, nil
	default:
		return 0, nil
	}
}
