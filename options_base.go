package gopium

import "strings"

func NewBaseOptions() *BaseOptions {
	return &BaseOptions{
		alwaysMatch:  map[string]any{},
		appiumOption: map[string]any{},
	}
}

func (o *BaseOptions) clone() *BaseOptions {
	return &BaseOptions{
		alwaysMatch:  cloneMap(o.alwaysMatch),
		firstMatch:   cloneMapSlice(o.firstMatch),
		appiumOption: cloneMap(o.appiumOption),
		useGrouping:  o.useGrouping,
	}
}

func (o *BaseOptions) WithAppiumOptions() *BaseOptions {
	o.useGrouping = true
	return o
}

func (o *BaseOptions) SetCapability(name string, value any) *BaseOptions {
	o.alwaysMatch[name] = value
	return o
}

func (o *BaseOptions) SetVendorOptions(name string, values map[string]any) *BaseOptions {
	name = strings.TrimSpace(name)
	if name == "" {
		return o
	}
	if !strings.Contains(name, ":") {
		name += ":options"
	}
	merged := map[string]any{}
	if existing, ok := o.alwaysMatch[name].(map[string]any); ok {
		merged = cloneMap(existing)
	}
	for key, value := range values {
		merged[key] = value
	}
	o.alwaysMatch[name] = merged
	return o
}

func (o *BaseOptions) SetCloudAppiumOptions(cloud string, values map[string]any) *BaseOptions {
	cloud = strings.TrimSpace(cloud)
	if cloud == "" {
		return o
	}
	name := cloud + ":appiumOptions"
	merged := map[string]any{}
	if existing, ok := o.alwaysMatch[name].(map[string]any); ok {
		merged = cloneMap(existing)
	}
	for key, value := range values {
		merged[key] = value
	}
	o.alwaysMatch[name] = merged
	return o
}

func (o *BaseOptions) SetAppiumCapability(name string, value any) *BaseOptions {
	name = strings.TrimSpace(name)
	if name == "" {
		return o
	}

	if strings.Contains(name, ":") {
		o.alwaysMatch[name] = value
		delete(o.appiumOption, strings.TrimPrefix(name, "appium:"))
		return o
	}

	if o.useGrouping {
		o.appiumOption[name] = value
		return o
	}

	o.alwaysMatch["appium:"+name] = value
	return o
}

func (o *BaseOptions) AddFirstMatch(capabilities map[string]any) *BaseOptions {
	o.firstMatch = append(o.firstMatch, cloneMap(capabilities))
	return o
}

func (o *BaseOptions) W3CCapabilities() W3CCapabilities {
	always := cloneMap(o.alwaysMatch)
	if o.useGrouping && len(o.appiumOption) > 0 {
		always["appium:options"] = cloneMap(o.appiumOption)
	}
	return W3CCapabilities{
		AlwaysMatch: always,
		FirstMatch:  cloneMapSlice(o.firstMatch),
	}
}

func (o *BaseOptions) LegacyCapabilities() map[string]any {
	legacy := map[string]any{}
	for key, value := range o.alwaysMatch {
		legacy[strings.TrimPrefix(key, "appium:")] = value
	}
	for key, value := range o.appiumOption {
		legacy[key] = value
	}
	return legacy
}

func (o *BaseOptions) IncludeLegacyCapabilities() bool {
	return false
}
