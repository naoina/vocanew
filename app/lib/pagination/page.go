package pagination

type Page struct {
	number int
	prev   *Page
	next   *Page
}

func (p *Page) Next() *Page {
	return p.next
}

func (p *Page) Prev() *Page {
	return p.prev
}

func (p *Page) HasNext() bool {
	return p.next != nil
}

func (p *Page) HasPrev() bool {
	return p.prev != nil
}

func (p *Page) Number() int {
	return p.number
}
