package gopium

import (
	"context"
	"net/http"
)

func (e *Element) FindElement(ctx context.Context, by By) (*Element, error) {
	envelope, err := e.driver.client.do(ctx, http.MethodPost, e.path("/element"), by.payload())
	if err != nil {
		return nil, err
	}

	id, err := envelope.decodeElementID()
	if err != nil {
		return nil, err
	}
	return &Element{driver: e.driver, id: id}, nil
}

func (e *Element) FindElements(ctx context.Context, by By) ([]*Element, error) {
	envelope, err := e.driver.client.do(ctx, http.MethodPost, e.path("/elements"), by.payload())
	if err != nil {
		return nil, err
	}

	ids, err := envelope.decodeElementIDs()
	if err != nil {
		return nil, err
	}

	out := make([]*Element, 0, len(ids))
	for _, id := range ids {
		out = append(out, &Element{driver: e.driver, id: id})
	}
	return out, nil
}
