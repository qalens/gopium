package gopium

import (
	"context"
	"net/http"
)

func (d *Driver) StartSession(ctx context.Context, options SessionOptions) error {
	if options == nil {
		options = NewBaseOptions()
	}

	envelope, err := d.client.do(ctx, http.MethodPost, "/session", newSessionRequest(options))
	if err != nil {
		return err
	}

	session, err := envelope.decodeSession()
	if err != nil {
		return err
	}

	d.sessionID = session.SessionID
	d.capabilities = session.Capabilities
	d.dialect = session.Dialect
	return nil
}

func (d *Driver) Quit(ctx context.Context) error {
	if d.sessionID == "" {
		return nil
	}

	_, err := d.client.do(ctx, http.MethodDelete, d.sessionPath(""), nil)
	if err != nil {
		return err
	}

	d.sessionID = ""
	d.capabilities = nil
	return nil
}

func (d *Driver) Status(ctx context.Context) (map[string]any, error) {
	return d.client.Status(ctx)
}

func (d *Driver) SessionID() string {
	return d.sessionID
}

func (d *Driver) Dialect() Dialect {
	return d.dialect
}

func (d *Driver) Capabilities() map[string]any {
	return cloneMap(d.capabilities)
}

func (d *Driver) GetSession(ctx context.Context) (map[string]any, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath(""), nil)
	if err != nil {
		return nil, err
	}
	return envelope.valueAsMap()
}

func (d *Driver) sessionPath(suffix string) string {
	return "/session/" + d.sessionID + suffix
}
