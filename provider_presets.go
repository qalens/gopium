package gopium

func NewBrowserStackProvider(username, accessKey string) *CloudProvider {
	return NewCloudProvider("BrowserStack", "https://hub-cloud.browserstack.com/wd/hub", "bstack:options").
		SetVendorOptions(map[string]any{
			"userName":  username,
			"accessKey": accessKey,
		}).
		SetCloudAppiumOptionsName("browserstack")
}

func NewSauceLabsProvider(username, accessKey string) *CloudProvider {
	return NewCloudProvider("Sauce Labs", "https://ondemand.us-west-1.saucelabs.com/wd/hub", "sauce:options").
		SetVendorOptions(map[string]any{
			"username":  username,
			"accessKey": accessKey,
		}).
		SetCloudAppiumOptionsName("sauce")
}

func NewLambdaTestProvider(username, accessKey string) *CloudProvider {
	return NewCloudProvider("LambdaTest", "https://mobile-hub.lambdatest.com/wd/hub", "LT:Options").
		WithBasicAuth(username, accessKey).
		SetCloudAppiumOptionsName("lambdaTest")
}

func NewTestingBotProvider(key, secret string) *CloudProvider {
	return NewCloudProvider("TestingBot", "https://hub.testingbot.com/wd/hub", "tb:options").
		SetVendorOptions(map[string]any{
			"key":    key,
			"secret": secret,
		}).
		SetCloudAppiumOptionsName("testingbot")
}
