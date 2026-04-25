package gopium

func NewAndroidOptions() *AndroidOptions {
	return &AndroidOptions{
		BaseOptions: NewBaseOptions().SetPlatformName("Android"),
	}
}

func (o *AndroidOptions) SetAutomationName(value string) *AndroidOptions {
	o.BaseOptions.SetAutomationName(value)
	return o
}

func (o *AndroidOptions) SetAppPackage(value string) *AndroidOptions {
	o.BaseOptions.SetAppPackage(value)
	return o
}

func (o *AndroidOptions) SetAppActivity(value string) *AndroidOptions {
	o.BaseOptions.SetAppActivity(value)
	return o
}

func (o *AndroidOptions) SetAppWaitActivity(value string) *AndroidOptions {
	o.BaseOptions.SetAppiumCapability("appWaitActivity", value)
	return o
}

func (o *AndroidOptions) SetAutoGrantPermissions(value bool) *AndroidOptions {
	o.BaseOptions.SetAutoGrantPermissions(value)
	return o
}

func NewUiAutomator2Options() *UiAutomator2Options {
	return &UiAutomator2Options{
		AndroidOptions: NewAndroidOptions().SetAutomationName("UiAutomator2"),
	}
}
