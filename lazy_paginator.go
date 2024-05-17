package pagination

import (
	"sync"
)

// LazyPaginator is a paginator with lazy load, it extends from [Paginator]
//
// Example:
//
//	var page = 1
//	var pageSize = 10
//	var tatolLoader = func() int64  {
//	    // fetch total
//	    c := db.Count()
//	    return c
//	}
//
//	var itemsLoader = func() []int {
//	    // fetch items
//	    var offset = (page-1)*pageSize
//	    var limit = pageSize
//	    items := db.Select("id").Offset(offset).Limit(limit).Find()
//	    return items
//	}
//
//	paginator := Lazy[int](tatolLoader, itemsLoader, pageSize, page)
type LazyPaginator[T any] struct {
	*Paginator[T]

	loader func()
	once   *sync.Once
}

// NewLazyPaginator creates a [LazyPaginator] instance
func NewLazyPaginator[T any](loader func() ([]T, int64), pageSize, page int) *LazyPaginator[T] {
	paginator := new(LazyPaginator[T])
	paginator.Paginator = new(Paginator[T])
	if page < 0 {
		page = 1
	}
	if pageSize < 0 {
		pageSize = 10
	}
	paginator.currentPage = page
	paginator.pageSize = pageSize
	paginator.loader = func() {
		paginator.items, paginator.total = loader()
		paginator.lastPage = int(paginator.total/int64(pageSize)) + 1
	}
	paginator.once = new(sync.Once)
	return paginator
}

func (p *LazyPaginator[T]) load() {
	p.once.Do(p.loader)
}

// Items returns all items of current page
func (p *LazyPaginator[T]) Items() []T {
	p.load()
	return p.Paginator.Items()
}

// LastItemIndex returns the index of the last item of current page
func (p *LazyPaginator[T]) LastItemIndex() int {
	p.load()
	return p.Paginator.LastItemIndex()
}

// Total returns the total count
func (p *LazyPaginator[T]) Total() int64 {
	p.load()
	return p.Paginator.Total()
}

// HasMore returns whether the next page exists
func (p *LazyPaginator[T]) HasMore() bool {
	return p.CurrentPage() < p.LastPage()
}

// LastPage returns the last page number
func (p *LazyPaginator[T]) LastPage() int {
	p.load()
	return p.Paginator.LastPage()
}

// ToMap returns current paginator info as a map
func (p *LazyPaginator[T]) ToMap() map[string]any {
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
