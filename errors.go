package gopium

import "fmt"

type Error struct {
	StatusCode int
	Code       string
	Message    string
	StackTrace string
}

func (e *Error) Error() string {
	switch {
	case e.Code != "" && e.Message != "":
		return fmt.Sprintf("%s: %s", e.Code, e.Message)
	case e.Message != "":
		return e.Message
	case e.Code != "":
		return e.Code
	default:
		return "appium request failed"
	}
}

func httpStatusFallback(statusCode int) string {
	if statusCode == 0 {
		return "request failed"
	}
	return fmt.Sprintf("request failed with status %d", statusCode)
}
