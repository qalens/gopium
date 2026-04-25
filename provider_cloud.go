package gopium

func NewCloudProvider(name, serverURL, vendorCapabilityName string) *CloudProvider {
	return &CloudProvider{
		name:                 name,
		serverURL:            serverURL,
		vendorCapabilityName: vendorCapabilityName,
		vendorOptions:        map[string]any{},
		cloudAppiumOptions:   map[string]any{},
	}
}

func (p *CloudProvider) Name() string {
	return p.name
}

func (p *CloudProvider) ServerURL() string {
	return p.serverURL
}

func (p *CloudProvider) ClientOptions() []ClientOption {
	out := make([]ClientOption, 0, len(p.clientOptions))
	out = append(out, p.clientOptions...)
	return out
}

func (p *CloudProvider) Transform(options SessionOptions) SessionOptions {
	base := NormalizeOptions(options)
	if p.vendorCapabilityName != "" && len(p.vendorOptions) > 0 {
		base.SetVendorOptions(p.vendorCapabilityName, p.vendorOptions)
	}
	if p.cloudAppiumName != "" && len(p.cloudAppiumOptions) > 0 {
		base.SetCloudAppiumOptions(p.cloudAppiumName, p.cloudAppiumOptions)
	}
	return base
}

func (p *CloudProvider) SetVendorOption(name string, value any) *CloudProvider {
	p.vendorOptions[name] = value
	return p
}

func (p *CloudProvider) SetVendorOptions(values map[string]any) *CloudProvider {
	for key, value := range values {
		p.vendorOptions[key] = value
	}
	return p
}

func (p *CloudProvider) SetCloudAppiumOptionsName(name string) *CloudProvider {
	p.cloudAppiumName = name
	return p
}

func (p *CloudProvider) SetCloudAppiumOption(name string, value any) *CloudProvider {
	p.cloudAppiumOptions[name] = value
	return p
}

func (p *CloudProvider) SetCloudAppiumOptions(values map[string]any) *CloudProvider {
	for key, value := range values {
		p.cloudAppiumOptions[key] = value
	}
	return p
}

func (p *CloudProvider) AddClientOptions(opts ...ClientOption) *CloudProvider {
	p.clientOptions = append(p.clientOptions, opts...)
	return p
}

func (p *CloudProvider) WithBasicAuth(username, password string) *CloudProvider {
	return p.AddClientOptions(WithBasicAuth(username, password))
}
