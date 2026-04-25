package gopium

type Driver struct {
	client       *Client
	sessionID    string
	dialect      Dialect
	capabilities map[string]any
}

type AppiumDriver = Driver
type AndroidDriver = Driver
type IOSDriver = Driver
