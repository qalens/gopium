package gopium

import (
	"context"
	"net/http"
)

func (d *Driver) ActiveElement(ctx context.Context) (*Element, error) {
	envelope, err := d.client.do(ctx, http.MethodGet, d.sessionPath("/element/active"), nil)
	if err != nil {
		return nil, err
	}

	id, err := envelope.decodeElementID()
	if err != nil {
		return nil, err
	}

	return &Element{driver: d, id: id}, nil
}

func (d *Driver) FindElement(ctx context.Context, by By) (*Element, error) {
	envelope, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/element"), by.payload())
	if err != nil {
		return nil, err
	}

	id, err := envelope.decodeElementID()
	if err != nil {
		return nil, err
	}

	return &Element{driver: d, id: id}, nil
}

func (d *Driver) FindElements(ctx context.Context, by By) ([]*Element, error) {
	envelope, err := d.client.do(ctx, http.MethodPost, d.sessionPath("/elements"), by.payload())
	if err != nil {
		return nil, err
	}

	ids, err := envelope.decodeElementIDs()
	if err != nil {
		return nil, err
	}

	elements := make([]*Element, 0, len(ids))
	for _, id := range ids {
		elements = append(elements, &Element{driver: d, id: id})
	}
	return elements, nil
}
