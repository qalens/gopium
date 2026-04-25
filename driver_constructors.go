package gopium

import "context"

func NewDriver(ctx context.Context, serverURL string, options SessionOptions, clientOpts ...ClientOption) (*Driver, error) {
	client, err := NewClient(serverURL, clientOpts...)
	if err != nil {
		return nil, err
	}
	return client.NewDriver(ctx, options)
}

func NewAppiumDriver(ctx context.Context, serverURL string, options SessionOptions, clientOpts ...ClientOption) (*AppiumDriver, error) {
	return NewDriver(ctx, serverURL, options, clientOpts...)
}

func NewAndroidDriver(ctx context.Context, serverURL string, options SessionOptions, clientOpts ...ClientOption) (*AndroidDriver, error) {
	return NewDriver(ctx, serverURL, options, clientOpts...)
}

func NewIOSDriver(ctx context.Context, serverURL string, options SessionOptions, clientOpts ...ClientOption) (*IOSDriver, error) {
	return NewDriver(ctx, serverURL, options, clientOpts...)
}
