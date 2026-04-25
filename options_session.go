package gopium

func newSessionRequest(options SessionOptions) newSessionPayload {
	w3c := options.W3CCapabilities()
	if len(w3c.FirstMatch) == 0 {
		w3c.FirstMatch = []map[string]any{{}}
	}

	return newSessionPayload{
		Capabilities:        w3c,
		DesiredCapabilities: options.LegacyCapabilities(),
	}
}

func NormalizeOptions(options SessionOptions) *BaseOptions {
	if options == nil {
		return NewBaseOptions()
	}
	if base, ok := options.(*BaseOptions); ok {
		return base.clone()
	}

	w3c := options.W3CCapabilities()
	normalized := NewBaseOptions()
	normalized.alwaysMatch = cloneMap(w3c.AlwaysMatch)
	normalized.firstMatch = cloneMapSlice(w3c.FirstMatch)
	return normalized
}
