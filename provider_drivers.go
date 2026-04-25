package gopium

import "context"

func NewRemoteDriver(ctx context.Context, provider Provider, options SessionOptions) (*Driver, error) {
	if provider == nil {
		return NewDriver(ctx, defaultServerURL, options)
	}
	return NewDriver(ctx, provider.ServerURL(), provider.Transform(options), provider.ClientOptions()...)
}

func NewProviderDriver(ctx context.Context, provider Provider, options SessionOptions) (*Driver, error) {
	return NewRemoteDriver(ctx, provider, options)
}
