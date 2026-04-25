package gopium

import (
	"encoding/json"
	"fmt"
)

func decodeResponse(statusCode int, data []byte) (*responseEnvelope, error) {
	if len(data) == 0 {
		return &responseEnvelope{}, nil
	}

	var envelope responseEnvelope
	if err := json.Unmarshal(data, &envelope); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if statusCode >= 400 {
		return nil, envelope.asError(statusCode)
	}
	if raw, ok := envelope.Value.(map[string]any); ok {
		if _, hasError := raw["error"]; hasError {
			return nil, envelope.asError(statusCode)
		}
	}
	if envelope.Status != nil && *envelope.Status != 0 {
		return nil, envelope.asError(statusCode)
	}

	return &envelope, nil
}

func (r *responseEnvelope) decodeSession() (sessionResult, error) {
	if value, ok := r.Value.(map[string]any); ok {
		if sessionID, ok := value["sessionId"].(string); ok {
			caps, _ := value["capabilities"].(map[string]any)
			return sessionResult{
				SessionID:    sessionID,
				Capabilities: cloneMap(caps),
				Dialect:      DialectW3C,
			}, nil
		}
		if r.SessionID != "" {
			return sessionResult{
				SessionID:    r.SessionID,
				Capabilities: cloneMap(value),
				Dialect:      DialectLegacy,
			}, nil
		}
	}

	if r.SessionID != "" {
		return sessionResult{
			SessionID:    r.SessionID,
			Capabilities: map[string]any{},
			Dialect:      DialectLegacy,
		}, nil
	}

	return sessionResult{}, fmt.Errorf("response did not contain a session id")
}

func (r *responseEnvelope) valueAsMap() (map[string]any, error) {
	value, ok := r.Value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("expected object response but got %T", r.Value)
	}
	return cloneMap(value), nil
}

func (r *responseEnvelope) decodeElementID() (string, error) {
	value, ok := r.Value.(map[string]any)
	if !ok {
		return "", fmt.Errorf("expected element object but got %T", r.Value)
	}
	return elementIDFromMap(value)
}

func (r *responseEnvelope) decodeElementIDs() ([]string, error) {
	items, ok := r.Value.([]any)
	if !ok {
		return nil, fmt.Errorf("expected element array but got %T", r.Value)
	}

	out := make([]string, 0, len(items))
	for _, item := range items {
		value, ok := item.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("expected element object but got %T", item)
		}
		id, err := elementIDFromMap(value)
		if err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, nil
}

func (r *responseEnvelope) asError(statusCode int) error {
	if value, ok := r.Value.(map[string]any); ok {
		var werr wireError
		if err := remarshal(value, &werr); err == nil && (werr.Error != "" || werr.Message != "") {
			return &Error{
				StatusCode: statusCode,
				Code:       werr.Error,
				Message:    werr.Message,
				StackTrace: werr.StackTrace,
			}
		}

		if message, ok := value["message"].(string); ok {
			return &Error{
				StatusCode: statusCode,
				Message:    message,
			}
		}
	}

	return &Error{
		StatusCode: statusCode,
		Message:    httpStatusFallback(statusCode),
	}
}

func elementIDFromMap(value map[string]any) (string, error) {
	if id, ok := value[w3cElementKey].(string); ok && id != "" {
		return id, nil
	}
	if id, ok := value[legacyElementKey].(string); ok && id != "" {
		return id, nil
	}
	return "", fmt.Errorf("response did not contain an element id")
}
