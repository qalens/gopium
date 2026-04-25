package gopium

func (o *BaseOptions) SetPlatformName(value string) *BaseOptions {
	return o.SetCapability("platformName", value)
}

func (o *BaseOptions) SetBrowserName(value string) *BaseOptions {
	return o.SetCapability("browserName", value)
}

func (o *BaseOptions) SetBrowserVersion(value string) *BaseOptions {
	return o.SetCapability("browserVersion", value)
}

func (o *BaseOptions) SetAutomationName(value string) *BaseOptions {
	return o.SetAppiumCapability("automationName", value)
}

func (o *BaseOptions) SetDeviceName(value string) *BaseOptions {
	return o.SetAppiumCapability("deviceName", value)
}

func (o *BaseOptions) SetPlatformVersion(value string) *BaseOptions {
	return o.SetAppiumCapability("platformVersion", value)
}

func (o *BaseOptions) SetUDID(value string) *BaseOptions {
	return o.SetAppiumCapability("udid", value)
}

func (o *BaseOptions) SetApp(value string) *BaseOptions {
	return o.SetAppiumCapability("app", value)
}

func (o *BaseOptions) SetBundleID(value string) *BaseOptions {
	return o.SetAppiumCapability("bundleId", value)
}

func (o *BaseOptions) SetAppPackage(value string) *BaseOptions {
	return o.SetAppiumCapability("appPackage", value)
}

func (o *BaseOptions) SetAppActivity(value string) *BaseOptions {
	return o.SetAppiumCapability("appActivity", value)
}

func (o *BaseOptions) SetNoReset(value bool) *BaseOptions {
	return o.SetAppiumCapability("noReset", value)
}

func (o *BaseOptions) SetFullReset(value bool) *BaseOptions {
	return o.SetAppiumCapability("fullReset", value)
}

func (o *BaseOptions) SetNewCommandTimeout(seconds int) *BaseOptions {
	return o.SetAppiumCapability("newCommandTimeout", seconds)
}

func (o *BaseOptions) SetLanguage(value string) *BaseOptions {
	return o.SetAppiumCapability("language", value)
}

func (o *BaseOptions) SetLocale(value string) *BaseOptions {
	return o.SetAppiumCapability("locale", value)
}

func (o *BaseOptions) SetAutoGrantPermissions(value bool) *BaseOptions {
	return o.SetAppiumCapability("autoGrantPermissions", value)
}

func (o *BaseOptions) SetWebSocketURL(value bool) *BaseOptions {
	return o.SetCapability("webSocketUrl", value)
}
