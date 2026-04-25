package gopium

type Provider interface {
	Name() string
	ServerURL() string
	ClientOptions() []ClientOption
	Transform(SessionOptions) SessionOptions
}

type CloudProvider struct {
	name                 string
	serverURL            string
	clientOptions        []ClientOption
	vendorCapabilityName string
	vendorOptions        map[string]any
	cloudAppiumName      string
	cloudAppiumOptions   map[string]any
}
