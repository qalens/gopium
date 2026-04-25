package gopium

type responseEnvelope struct {
	SessionID string `json:"sessionId,omitempty"`
	Status    *int   `json:"status,omitempty"`
	Value     any    `json:"value"`
}

type sessionResult struct {
	SessionID    string
	Capabilities map[string]any
	Dialect      Dialect
}

type wireError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StackTrace string `json:"stacktrace"`
}
