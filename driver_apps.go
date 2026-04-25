package gopium

import (
	"context"
	"net/http"
)

func (d *Driver) LaunchApp(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/app/launch"), map[string]any{})
	return err
}

func (d *Driver) CloseApp(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/app/close"), map[string]any{})
	return err
}

func (d *Driver) ResetApp(ctx context.Context) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/app/reset"), map[string]any{})
	return err
}

func (d *Driver) BackgroundApp(ctx context.Context, seconds int) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/app/background"), map[string]any{"seconds": seconds})
	return err
}

func (d *Driver) InstallApp(ctx context.Context, appPath string, options InstallAppOptions) error {
	payload := map[string]any{"appPath": appPath}
	if len(options) > 0 {
		payload["options"] = map[string]any(options)
	}
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/install_app"), payload)
	return err
}

func (d *Driver) RemoveApp(ctx context.Context, appID string, options RemoveAppOptions) error {
	payload := map[string]any{"appId": appID}
	if len(options) > 0 {
		payload["options"] = map[string]any(options)
	}
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/remove_app"), payload)
	return err
}

func (d *Driver) ActivateApp(ctx context.Context, appID string) error {
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/activate_app"), map[string]any{"appId": appID})
	return err
}

func (d *Driver) TerminateApp(ctx context.Context, appID string, options TerminateAppOptions) (bool, error) {
	payload := map[string]any{"appId": appID}
	if len(options) > 0 {
		payload["options"] = map[string]any(options)
	}
	return d.postBool(ctx, "/appium/device/terminate_app", payload)
}

func (d *Driver) AppState(ctx context.Context, appID string) (AppState, error) {
	envelope, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/app_state"), map[string]any{"appId": appID})
	if err != nil {
		return AppStateNotInstalled, err
	}
	switch v := envelope.Value.(type) {
	case float64:
		return AppState(int(v)), nil
	case int:
		return AppState(v), nil
	default:
		return AppStateNotInstalled, nil
	}
}

func (d *Driver) StartActivity(ctx context.Context, activity AppActivity) error {
	payload := map[string]any{
		"appPackage":  activity.AppPackage,
		"appActivity": activity.AppActivity,
	}
	if activity.AppWaitPackage != "" {
		payload["appWaitPackage"] = activity.AppWaitPackage
	}
	if activity.AppWaitActivity != "" {
		payload["appWaitActivity"] = activity.AppWaitActivity
	}
	if activity.IntentAction != "" {
		payload["intentAction"] = activity.IntentAction
	}
	if activity.IntentCategory != "" {
		payload["intentCategory"] = activity.IntentCategory
	}
	if activity.IntentFlags != "" {
		payload["intentFlags"] = activity.IntentFlags
	}
	if activity.OptionalIntentArguments != "" {
		payload["optionalIntentArguments"] = activity.OptionalIntentArguments
	}
	if activity.StopApp {
		payload["stopApp"] = true
	}
	_, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/appium/device/start_activity"), payload)
	return err
}
