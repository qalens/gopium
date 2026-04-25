package gopium

import (
	"context"
	"net/http"
)

func (e *Element) Click(ctx context.Context) error {
	_, err := e.driver.client.do(ctx, http.MethodPost, e.path("/click"), map[string]any{})
	return err
}

func (e *Element) Clear(ctx context.Context) error {
	_, err := e.driver.client.do(ctx, http.MethodPost, e.path("/clear"), map[string]any{})
	return err
}

func (e *Element) SendKeys(ctx context.Context, text string) error {
	_, err := e.driver.client.do(ctx, http.MethodPost, e.path("/value"), map[string]any{
		"text":  text,
		"value": splitRunes(text),
	})
	return err
}
