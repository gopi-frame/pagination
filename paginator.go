package pagination

import (
	"github.com/gopi-frame/contract/pagination"
)

var _ pagination.Paginator[any] = (*Paginator[any])(nil)
var _ pagination.Paginator[any] = (*ArrayPaginator[any])(nil)
var _ pagination.Paginator[any] = (*LazyPaginator[any])(nil)

// Paginator is the basic paginator, it implements [IPaginator] interface
//
// Example:
//
//	total := 100
//	items := []int{1,2,3,4,5,6,7,8,9,10}
//	pageSize := 10
//	page := 1
//	paginator := New[int](total, items, pageSize, page)
type Paginator[T any] struct {
	items       []T
	total       int64
	currentPage int
	pageSize    int
	lastPage    int
}

// NewPaginator creates a [Paginator] instance
func NewPaginator[T any](total int64, items []T, pageSize, page int) *Paginator[T] {
	if page < 0 {
		page = 1
	}
	if pageSize < 0 {
		pageSize = 10
	}
	paginator := new(Paginator[T])
	paginator.items = items
	paginator.total = total
	paginator.currentPage = page
	paginator.pageSize = pageSize
	paginator.lastPage = int((total / int64(pageSize))) + 1
	return paginator
}

// Items returns all items of current page
func (p *Paginator[T]) Items() []T {
	return p.items
}

// Total returns the total count
func (p *Paginator[T]) Total() int64 {
	return p.total
}

// FirstItemIndex returns the index of the first item of current page
func (p *Paginator[T]) FirstItemIndex() int {
	return (p.currentPage - 1) * p.pageSize
}

// LastItemIndex returns the index of the last item of current page
func (p *Paginator[T]) LastItemIndex() int {
	return p.FirstItemIndex() + len(p.items) - 1
}

// CurrentPage returns the current page number
func (p *Paginator[T]) CurrentPage() int {
	return p.currentPage
}

// PageSize returns the page size
func (p *Paginator[T]) PageSize() int {
	return p.pageSize
}

// HasMore returns whether the next page exists
func (p *Paginator[T]) HasMore() bool {
	return p.lastPage > p.currentPage
}

// LastPage returns the last page number
func (p *Paginator[T]) LastPage() int {
	return p.lastPage
}

// ToMap returns current paginator info as a map
func (p *Paginator[T]) ToMap() map[string]any {
	return map[string]any{
		"items":        p.Items(),
		"total":        p.Total(),
		"current_page": p.CurrentPage(),
		"page_size":    p.PageSize(),
		"last_page":    p.LastPage(),
		"from":         p.FirstItemIndex(),
		"to":           p.LastItemIndex(),
	}
}
