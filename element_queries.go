package gopium

import (
	"context"
	"fmt"
	"net/http"
)

func (e *Element) Text(ctx context.Context) (string, error) {
	return e.readString(ctx, "/text")
}

func (e *Element) Name(ctx context.Context) (string, error) {
	return e.readString(ctx, "/name")
}

func (e *Element) Attribute(ctx context.Context, name string) (string, error) {
	return e.readString(ctx, "/attribute/"+name)
}

func (e *Element) Property(ctx context.Context, name string) (any, error) {
	envelope, err := e.driver.client.do(ctx, http.MethodGet, e.path("/property/"+name), nil)
	if err != nil {
		return nil, err
	}
	return envelope.Value, nil
}

func (e *Element) CSSValue(ctx context.Context, name string) (string, error) {
	return e.readString(ctx, "/css/"+name)
}

func (e *Element) Rect(ctx context.Context) (Rect, error) {
	envelope, err := e.driver.client.do(ctx, http.MethodGet, e.path("/rect"), nil)
	if err != nil {
		return Rect{}, err
	}

	var rect Rect
	if err := remarshal(envelope.Value, &rect); err != nil {
		return Rect{}, fmt.Errorf("decode rect: %w", err)
	}
	return rect, nil
}

func (e *Element) Screenshot(ctx context.Context) (string, error) {
	return e.readString(ctx, "/screenshot")
}

func (e *Element) IsDisplayed(ctx context.Context) (bool, error) {
	return e.readBool(ctx, "/displayed")
}

func (e *Element) IsEnabled(ctx context.Context) (bool, error) {
	return e.readBool(ctx, "/enabled")
}

func (e *Element) IsSelected(ctx context.Context) (bool, error) {
	return e.readBool(ctx, "/selected")
}

func (e *Element) readString(ctx context.Context, suffix string) (string, error) {
	envelope, err := e.driver.client.do(ctx, http.MethodGet, e.path(suffix), nil)
	if err != nil {
		return "", err
	}

	value, ok := envelope.Value.(string)
	if !ok {
		return "", fmt.Errorf("expected string response but got %T", envelope.Value)
	}
	return value, nil
}

func (e *Element) readBool(ctx context.Context, suffix string) (bool, error) {
	envelope, err := e.driver.client.do(ctx, http.MethodGet, e.path(suffix), nil)
	if err != nil {
		return false, err
	}

	value, ok := envelope.Value.(bool)
	if !ok {
		return false, fmt.Errorf("expected bool response but got %T", envelope.Value)
	}
	return value, nil
}
