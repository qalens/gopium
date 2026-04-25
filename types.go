package gopium

type Rect struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Timeouts struct {
	Implicit float64 `json:"implicit,omitempty"`
	PageLoad float64 `json:"pageLoad,omitempty"`
	Script   float64 `json:"script,omitempty"`
}

type Rotation struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
	HTTPOnly bool   `json:"httpOnly,omitempty"`
	Expiry   int64  `json:"expiry,omitempty"`
	SameSite string `json:"sameSite,omitempty"`
}

type Action struct {
	Type       string         `json:"type"`
	ID         string         `json:"id,omitempty"`
	Parameters map[string]any `json:"parameters,omitempty"`
	Actions    []ActionItem   `json:"actions,omitempty"`
}

type ActionItem struct {
	Type     string         `json:"type"`
	Duration int64          `json:"duration,omitempty"`
	X        int            `json:"x,omitempty"`
	Y        int            `json:"y,omitempty"`
	Button   int            `json:"button,omitempty"`
	Value    string         `json:"value,omitempty"`
	Origin   any            `json:"origin,omitempty"`
	Extra    map[string]any `json:"-"`
}

type ClipboardContentType string

const (
	ClipboardPlaintext ClipboardContentType = "plaintext"
	ClipboardImage     ClipboardContentType = "image"
	ClipboardURL       ClipboardContentType = "url"
)

type AppState int

const (
	AppStateNotInstalled AppState = iota
	AppStateNotRunning
	AppStateRunningInBackgroundSuspended
	AppStateRunningInBackground
	AppStateRunningInForeground
)

type AppActivity struct {
	AppPackage              string
	AppActivity             string
	AppWaitPackage          string
	AppWaitActivity         string
	IntentAction            string
	IntentCategory          string
	IntentFlags             string
	OptionalIntentArguments string
	StopApp                 bool
}

type InstallAppOptions map[string]any
type RemoveAppOptions map[string]any
type TerminateAppOptions map[string]any
type RecordScreenOptions map[string]any
