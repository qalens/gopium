package gopium

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) do(ctx context.Context, method, path string, body any) (*responseEnvelope, error) {
	var payload io.Reader
	var payloadBytes []byte
	if body != nil {
		buf := &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, fmt.Errorf("encode request body: %w", err)
		}
		payloadBytes = append(payloadBytes, buf.Bytes()...)
		payload = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, c.resolve(path), payload)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("User-Agent", c.userAgent)
	for _, editor := range c.editors {
		editor(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s %s: %w", method, req.URL, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	if c.logger != nil {
		c.logger.Printf("[gopium] %s %s request=%s response_status=%d response=%s",
			method,
			req.URL.String(),
			string(payloadBytes),
			resp.StatusCode,
			string(data),
		)
	}

	return decodeResponse(resp.StatusCode, data)
}
