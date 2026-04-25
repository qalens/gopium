package gopium

func NewIOSOptions() *IOSOptions {
	return &IOSOptions{
		BaseOptions: NewBaseOptions().SetPlatformName("iOS"),
	}
}

func (o *IOSOptions) SetAutomationName(value string) *IOSOptions {
	o.BaseOptions.SetAutomationName(value)
	return o
}

func (o *IOSOptions) SetBundleID(value string) *IOSOptions {
	o.BaseOptions.SetBundleID(value)
	return o
}

func (o *IOSOptions) SetWdaLocalPort(value int) *IOSOptions {
	o.BaseOptions.SetAppiumCapability("wdaLocalPort", value)
	return o
}

func NewXCUITestOptions() *XCUITestOptions {
	return &XCUITestOptions{
		IOSOptions: NewIOSOptions().SetAutomationName("XCUITest"),
	}
}
