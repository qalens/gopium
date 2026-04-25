package gopium

type SessionOptions interface {
	W3CCapabilities() W3CCapabilities
	LegacyCapabilities() map[string]any
}

type W3CCapabilities struct {
	AlwaysMatch map[string]any   `json:"alwaysMatch,omitempty"`
	FirstMatch  []map[string]any `json:"firstMatch,omitempty"`
}

type newSessionPayload struct {
	Capabilities        W3CCapabilities `json:"capabilities"`
	DesiredCapabilities map[string]any  `json:"desiredCapabilities,omitempty"`
}

type BaseOptions struct {
	alwaysMatch  map[string]any
	firstMatch   []map[string]any
	appiumOption map[string]any
	useGrouping  bool
}

type AndroidOptions struct {
	*BaseOptions
}

type IOSOptions struct {
	*BaseOptions
}

type UiAutomator2Options struct {
	*AndroidOptions
}

type XCUITestOptions struct {
	*IOSOptions
}
