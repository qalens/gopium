package gopium

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"qalens.com/gopium/internal/testsupport"
)

func TestNewSessionRequestIncludesW3CAndLegacyPayloads(t *testing.T) {
	options := NewUiAutomator2Options().
		SetDeviceName("Pixel 9").
		SetUDID("emulator-5554").
		SetApp("/tmp/demo.apk")

	req := newSessionRequest(options)

	if got := req.Capabilities.AlwaysMatch["platformName"]; got != "Android" {
		t.Fatalf("platformName = %v", got)
	}
	if got := req.Capabilities.AlwaysMatch["appium:automationName"]; got != "UiAutomator2" {
		t.Fatalf("automationName = %v", got)
	}
	if got := req.DesiredCapabilities["automationName"]; got != "UiAutomator2" {
		t.Fatalf("legacy automationName = %v", got)
	}
	if len(req.Capabilities.FirstMatch) != 1 {
		t.Fatalf("firstMatch len = %d", len(req.Capabilities.FirstMatch))
	}
}

func TestBaseOptionsCanGroupAppiumOptions(t *testing.T) {
	options := NewBaseOptions().
		WithAppiumOptions().
		SetPlatformName("iOS").
		SetAutomationName("XCUITest").
		SetDeviceName("iPhone 16")

	w3c := options.W3CCapabilities()
	grouped, ok := w3c.AlwaysMatch["appium:options"].(map[string]any)
	if !ok {
		t.Fatalf("expected grouped appium:options")
	}
	if _, found := w3c.AlwaysMatch["appium:automationName"]; found {
		t.Fatalf("did not expect flattened appium capability when grouping")
	}
	if got := grouped["automationName"]; got != "XCUITest" {
		t.Fatalf("grouped automationName = %v", got)
	}
	if got := options.LegacyCapabilities()["automationName"]; got != "XCUITest" {
		t.Fatalf("legacy grouped automationName = %v", got)
	}
}

func TestDriverStartSessionSupportsW3CResponse(t *testing.T) {
	client, err := NewClient("http://127.0.0.1:4723", WithHTTPClient(testsupport.RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/session" {
			t.Fatalf("unexpected path %s", req.URL.Path)
		}

		body, _ := io.ReadAll(req.Body)
		if !strings.Contains(string(body), `"capabilities"`) {
			t.Fatalf("missing capabilities payload: %s", string(body))
		}

		return testsupport.JSONResponse(http.StatusOK, `{"value":{"sessionId":"w3c-session","capabilities":{"platformName":"Android"}}}`), nil
	})))
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	driver, err := client.NewDriver(context.Background(), NewBaseOptions().SetPlatformName("Android"))
	if err != nil {
		t.Fatalf("NewDriver error: %v", err)
	}

	if driver.SessionID() != "w3c-session" {
		t.Fatalf("session id = %s", driver.SessionID())
	}
	if driver.Dialect() != DialectW3C {
		t.Fatalf("dialect = %s", driver.Dialect())
	}
}

func TestDriverStartSessionSupportsLegacyResponse(t *testing.T) {
	client, err := NewClient("http://127.0.0.1:4723", WithHTTPClient(testsupport.RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		return testsupport.JSONResponse(http.StatusOK, `{"status":0,"sessionId":"legacy-session","value":{"platform":"Android"}}`), nil
	})))
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	driver, err := client.NewDriver(context.Background(), NewBaseOptions().SetPlatformName("Android"))
	if err != nil {
		t.Fatalf("NewDriver error: %v", err)
	}

	if driver.SessionID() != "legacy-session" {
		t.Fatalf("session id = %s", driver.SessionID())
	}
	if driver.Dialect() != DialectLegacy {
		t.Fatalf("dialect = %s", driver.Dialect())
	}
}

func TestResponseEnvelopeDecodesBothElementFormats(t *testing.T) {
	cases := []string{
		`{"value":{"element-6066-11e4-a52e-4f735466cecf":"w3c"}}`,
		`{"value":{"ELEMENT":"legacy"}}`,
	}

	for _, tc := range cases {
		var envelope responseEnvelope
		if err := json.Unmarshal([]byte(tc), &envelope); err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		id, err := envelope.decodeElementID()
		if err != nil {
			t.Fatalf("decodeElementID error: %v", err)
		}
		if id == "" {
			t.Fatalf("decoded empty id")
		}
	}
}

func TestDecodeResponseReturnsWireError(t *testing.T) {
	_, err := decodeResponse(http.StatusNotFound, []byte(`{"value":{"error":"no such element","message":"not found"}}`))
	if err == nil {
		t.Fatal("expected error")
	}

	appiumErr, ok := err.(*Error)
	if !ok {
		t.Fatalf("expected *Error, got %T", err)
	}
	if appiumErr.Code != "no such element" {
		t.Fatalf("error code = %s", appiumErr.Code)
	}
}

func TestCloudProviderTransformsOptions(t *testing.T) {
	provider := NewBrowserStackProvider("user", "key").
		SetVendorOption("projectName", "demo").
		SetCloudAppiumOption("version", "2.15.0")

	options := provider.Transform(NewUiAutomator2Options().
		SetDeviceName("Pixel 9").
		SetApp("bs://app"))

	w3c := options.W3CCapabilities()
	bstack, ok := w3c.AlwaysMatch["bstack:options"].(map[string]any)
	if !ok {
		t.Fatal("expected bstack:options")
	}
	if bstack["userName"] != "user" || bstack["accessKey"] != "key" {
		t.Fatalf("unexpected BrowserStack credentials: %#v", bstack)
	}
	if got := w3c.AlwaysMatch["browserstack:appiumOptions"]; got == nil {
		t.Fatal("expected browserstack:appiumOptions")
	}
}

func TestNewNamedProviderSupportsKnownClouds(t *testing.T) {
	tests := []struct {
		name           string
		provider       string
		creds          map[string]any
		wantName       string
		wantURL        string
		wantOpt        string
		wantClientOpts int
	}{
		{
			name:     "browserstack aliases",
			provider: "BSACCOUNT",
			creds: map[string]any{
				"username":   "alice",
				"access_key": "secret",
			},
			wantName: "BrowserStack",
			wantURL:  "https://hub-cloud.browserstack.com/wd/hub",
			wantOpt:  "bstack:options",
		},
		{
			name:     "saucelabs aliases",
			provider: "sauce",
			creds: map[string]any{
				"userName":  "alice",
				"accessKey": "secret",
			},
			wantName: "Sauce Labs",
			wantURL:  "https://ondemand.us-west-1.saucelabs.com/wd/hub",
			wantOpt:  "sauce:options",
		},
		{
			name:     "lambdatest aliases",
			provider: "lambda-test",
			creds: map[string]any{
				"username": "alice",
				"key":      "secret",
			},
			wantName:       "LambdaTest",
			wantURL:        "https://mobile-hub.lambdatest.com/wd/hub",
			wantClientOpts: 1,
		},
		{
			name:     "testingbot aliases",
			provider: "testing_bot",
			creds: map[string]any{
				"key":    "alice",
				"secret": "secret",
			},
			wantName: "TestingBot",
			wantURL:  "https://hub.testingbot.com/wd/hub",
			wantOpt:  "tb:options",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			provider, err := NewNamedProvider(tc.provider, tc.creds)
			if err != nil {
				t.Fatalf("NewNamedProvider() error = %v", err)
			}
			if provider.Name() != tc.wantName {
				t.Fatalf("provider name = %q, want %q", provider.Name(), tc.wantName)
			}
			if provider.ServerURL() != tc.wantURL {
				t.Fatalf("provider url = %q, want %q", provider.ServerURL(), tc.wantURL)
			}
			if got := len(provider.ClientOptions()); got != tc.wantClientOpts {
				t.Fatalf("client options len = %d, want %d", got, tc.wantClientOpts)
			}

			options := provider.Transform(NewBaseOptions().SetPlatformName("Android"))
			caps := options.W3CCapabilities().AlwaysMatch
			if tc.wantOpt != "" {
				if _, ok := caps[tc.wantOpt]; !ok {
					t.Fatalf("expected %s in capabilities", tc.wantOpt)
				}
			}
		})
	}
}

func TestNewNamedProviderRejectsUnknownProvider(t *testing.T) {
	if _, err := NewNamedProvider("unknown-cloud", map[string]any{}); err == nil {
		t.Fatal("expected error for unknown provider")
	}
}

func TestWithBasicAuthAddsAuthorizationHeader(t *testing.T) {
	client, err := NewClient("https://example.com", WithBasicAuth("alice", "secret"), WithHTTPClient(testsupport.RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		user, pass, ok := req.BasicAuth()
		if !ok || user != "alice" || pass != "secret" {
			t.Fatalf("missing basic auth header")
		}
		return testsupport.JSONResponse(http.StatusOK, `{"value":{}}`), nil
	})))
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	if _, err := client.Status(context.Background()); err != nil {
		t.Fatalf("Status error: %v", err)
	}
}

func TestRecordingAndClipboardCommandsUseAppiumRoutes(t *testing.T) {
	var paths []string
	client, err := NewClient("http://127.0.0.1:4723", WithHTTPClient(testsupport.RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		paths = append(paths, req.URL.Path)
		body, _ := io.ReadAll(req.Body)
		switch req.URL.Path {
		case "/session/test-session/appium/start_recording_screen":
			if !strings.Contains(string(body), `"timeLimit":"180"`) {
				t.Fatalf("unexpected start recording payload: %s", body)
			}
			return testsupport.JSONResponse(http.StatusOK, `{"value":null}`), nil
		case "/session/test-session/appium/stop_recording_screen":
			return testsupport.JSONResponse(http.StatusOK, `{"value":"video-base64"}`), nil
		case "/session/test-session/appium/device/get_clipboard":
			return testsupport.JSONResponse(http.StatusOK, `{"value":"`+base64.StdEncoding.EncodeToString([]byte("hello"))+`"}`), nil
		case "/session/test-session/appium/device/set_clipboard":
			if !strings.Contains(string(body), `"contentType":"plaintext"`) {
				t.Fatalf("unexpected clipboard payload: %s", body)
			}
			return testsupport.JSONResponse(http.StatusOK, `{"value":null}`), nil
		default:
			t.Fatalf("unexpected path %s", req.URL.Path)
			return nil, nil
		}
	})))
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	driver := &Driver{client: client, sessionID: "test-session"}
	if err := driver.StartRecordingScreen(context.Background(), RecordScreenOptions{"timeLimit": "180"}); err != nil {
		t.Fatalf("StartRecordingScreen error: %v", err)
	}
	video, err := driver.StopRecordingScreen(context.Background(), nil)
	if err != nil || video != "video-base64" {
		t.Fatalf("StopRecordingScreen = %q, %v", video, err)
	}
	text, err := driver.GetClipboardText(context.Background())
	if err != nil || text != "hello" {
		t.Fatalf("GetClipboardText = %q, %v", text, err)
	}
	if err := driver.SetClipboardText(context.Background(), "greeting", "hello"); err != nil {
		t.Fatalf("SetClipboardText error: %v", err)
	}
	if len(paths) != 4 {
		t.Fatalf("unexpected calls: %v", paths)
	}
}

func TestAppLifecycleAndContextsCommands(t *testing.T) {
	client, err := NewClient("http://127.0.0.1:4723", WithHTTPClient(testsupport.RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		switch req.URL.Path {
		case "/session/test-session/contexts":
			return testsupport.JSONResponse(http.StatusOK, `{"value":["NATIVE_APP","WEBVIEW_1"]}`), nil
		case "/session/test-session/context":
			if req.Method == http.MethodGet {
				return testsupport.JSONResponse(http.StatusOK, `{"value":"NATIVE_APP"}`), nil
			}
			return testsupport.JSONResponse(http.StatusOK, `{"value":null}`), nil
		case "/session/test-session/appium/device/terminate_app":
			return testsupport.JSONResponse(http.StatusOK, `{"value":true}`), nil
		case "/session/test-session/appium/device/app_state":
			return testsupport.JSONResponse(http.StatusOK, `{"value":4}`), nil
		case "/session/test-session/actions":
			return testsupport.JSONResponse(http.StatusOK, `{"value":null}`), nil
		default:
			t.Fatalf("unexpected path %s", req.URL.Path)
			return nil, nil
		}
	})))
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	driver := &Driver{client: client, sessionID: "test-session"}
	contexts, err := driver.Contexts(context.Background())
	if err != nil || len(contexts) != 2 {
		t.Fatalf("Contexts = %v, %v", contexts, err)
	}
	current, err := driver.CurrentContext(context.Background())
	if err != nil || current != "NATIVE_APP" {
		t.Fatalf("CurrentContext = %q, %v", current, err)
	}
	if err := driver.SetContext(context.Background(), "WEBVIEW_1"); err != nil {
		t.Fatalf("SetContext error: %v", err)
	}
	ok, err := driver.TerminateApp(context.Background(), "com.demo", nil)
	if err != nil || !ok {
		t.Fatalf("TerminateApp = %v, %v", ok, err)
	}
	state, err := driver.AppState(context.Background(), "com.demo")
	if err != nil || state != AppStateRunningInForeground {
		t.Fatalf("AppState = %v, %v", state, err)
	}
	if err := driver.PerformActions(context.Background(), []Action{{Type: "pointer"}}); err != nil {
		t.Fatalf("PerformActions error: %v", err)
	}
	if err := driver.ReleaseActions(context.Background()); err != nil {
		t.Fatalf("ReleaseActions error: %v", err)
	}
}

func TestSessionCommandEscapeHatch(t *testing.T) {
	client, err := NewClient("http://127.0.0.1:4723", WithHTTPClient(testsupport.RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		switch req.URL.Path {
		case "/session/test-session/appium/custom_extension":
			return testsupport.JSONResponse(http.StatusOK, `{"value":{"ok":true}}`), nil
		case "/session/test-session/goog/cdp/execute":
			return testsupport.JSONResponse(http.StatusOK, `{"value":{"result":"done"}}`), nil
		default:
			t.Fatalf("unexpected path %s", req.URL.Path)
			return nil, nil
		}
	})))
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	driver := &Driver{client: client, sessionID: "test-session"}
	value, err := driver.AppiumCommand(context.Background(), "custom_extension", map[string]any{"x": 1})
	if err != nil {
		t.Fatalf("AppiumCommand error: %v", err)
	}
	if value == nil {
		t.Fatal("expected AppiumCommand result")
	}
	value, err = driver.ExecuteCDP(context.Background(), "goog", "Network.enable", nil)
	if err != nil {
		t.Fatalf("ExecuteCDP error: %v", err)
	}
	if value == nil {
		t.Fatal("expected ExecuteCDP result")
	}
}
