package gopium

type Element struct {
	driver *Driver
	id     string
}

func (e *Element) ID() string {
	return e.id
}

func (e *Element) path(suffix string) string {
	return e.driver.sessionPath("/element/" + e.id + suffix)
}
