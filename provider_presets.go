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
	return NewSauceLabsProviderForRegion(username, accessKey, "")
}

func NewSauceLabsProviderForRegion(username, accessKey, region string) *CloudProvider {
	return NewCloudProvider("Sauce Labs", sauceLabsServerURL(region), "sauce:options").
		WithBasicAuth(username, accessKey).
		SetVendorOptions(map[string]any{
			"username":      username,
			"accessKey":     accessKey,
			"appiumVersion": "latest",
		}).
		SetCloudAppiumOptionsName("sauce")
}

func sauceLabsServerURL(region string) string {
	switch normalizeProviderName(region) {
	case "", "uswest", "uswest1", "us":
		return "https://ondemand.us-west-1.saucelabs.com/wd/hub"
	case "useast", "useast4":
		return "https://ondemand.us-east-4.saucelabs.com/wd/hub"
	case "eu", "eucentral", "eucentral1":
		return "https://ondemand.eu-central-1.saucelabs.com/wd/hub"
	default:
		return "https://ondemand.us-west-1.saucelabs.com/wd/hub"
	}
}

func NewBitBarProvider(apiKey string) *CloudProvider {
	return NewBitBarProviderForRegion(apiKey, "")
}

func NewBitBarProviderForRegion(apiKey, region string) *CloudProvider {
	return NewCloudProvider("BitBar", bitBarServerURL(region), "bitbar:options").
		SetVendorOptions(map[string]any{
			"apiKey": apiKey,
		})
}

func bitBarServerURL(region string) string {
	switch normalizeProviderName(region) {
	case "", "us", "uswest", "uswestmobile", "uswest1":
		return "https://us-west-mobile-hub.bitbar.com/wd/hub"
	case "eu", "euwest", "eumobile", "eucentral", "eucentral1":
		return "https://eu-mobile-hub.bitbar.com/wd/hub"
	default:
		return "https://us-west-mobile-hub.bitbar.com/wd/hub"
	}
}

func NewLambdaTestProvider(username, accessKey string) *CloudProvider {
	return NewCloudProvider("LambdaTest", "https://mobile-hub.lambdatest.com/wd/hub", "LT:Options").
		WithBasicAuth(username, accessKey).
		SetCloudAppiumOptionsName("lambdaTest")
}

func NewTestingBotProvider(key, secret string) *CloudProvider {
	return NewCloudProvider("TestingBot", "https://hub.testingbot.com/wd/hub", "tb:options").
		WithBasicAuth(key, secret).
		SetVendorOptions(map[string]any{
			"key":    key,
			"secret": secret,
		}).
		SetCloudAppiumOptionsName("testingbot")
}
