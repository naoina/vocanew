package pagination

import (
	"math"
)

type Paginator struct {
	pages []*Page

	// Number of all objects
	Count int64

	// Number of pages
	NumPages int

	// Items per page
	perPage int

	paginationLength int
	threshold        int
}

func NewPaginator(length int64, perPage, paginationLength, threshold int) *Paginator {
	numPages, r := int(length/int64(perPage)), length%int64(perPage)
	if r > 0 {
		numPages++
	}
	numPages = int(math.Max(float64(numPages), 1))
	paginator := &Paginator{Count: length, NumPages: numPages, perPage: perPage, paginationLength: paginationLength, threshold: threshold}
	return paginator
}

func (p *Paginator) Page(page int) *Page {
	page = p.round(page)
	pages := p.genPages(page)
	for _, v := range pages {
		if v.Number() == page {
			return v
		}
	}
	return pages[0]
}

func (p *Paginator) Pages(page int) []*Page {
	return p.genPages(p.round(page))
}

func (p *Paginator) round(x int) int {
	return int(math.Min(math.Max(float64(x), 1), float64(p.NumPages)))
}

func (p *Paginator) genPages(page int) []*Page {
	if p.pages != nil {
		return p.pages
	}
	pages := make([]*Page, 0, p.paginationLength)
	var base int
	if page <= p.threshold {
		base = 0
	} else if p.NumPages-p.threshold <= page {
		base = p.NumPages - p.paginationLength
	} else {
		base = page - p.threshold
	}
	for i := 1; i <= p.paginationLength; i++ {
		pages = append(pages, &Page{number: base + i})
	}
	for i, v := range pages {
		prevIndex := i - 1
		if prevIndex >= 0 {
			v.prev = pages[prevIndex]
		}
		nextIndex := i + 1
		if nextIndex < p.paginationLength {
			v.next = pages[nextIndex]
		}
	}
	p.pages = pages
	return pages
}
