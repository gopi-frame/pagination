package pagination

import (
	"github.com/gopi-frame/exception"
	"sync"
)

type LazyPaginator[T any] struct {
	Paginator[T]
	once     sync.Once
	loader   func(page, pageSize int) ([]T, int64)
	page     int
	pageSize int
}

func NewLazy[T any](loader func(page, pageSize int) ([]T, int64), page int, pageSize int) *LazyPaginator[T] {
	page = max(1, page)
	if pageSize <= 0 {
		pageSize = 10
	}
	return &LazyPaginator[T]{
		loader:   loader,
		page:     page,
		pageSize: pageSize,
	}
}

func (p *LazyPaginator[T]) init() {
	p.once.Do(func() {
		items, total := p.loader(p.page, p.pageSize)
		p.Paginator = *New[T](items, total, p.page, p.pageSize)
	})
}

func (p *LazyPaginator[T]) Items() []T {
	p.init()
	return p.Paginator.Items()
}

func (p *LazyPaginator[T]) Total() int64 {
	p.init()
	return p.Paginator.Total()
}

func (p *LazyPaginator[T]) LastPage() int {
	p.init()
	return p.Paginator.LastPage()
}

func (p *LazyPaginator[T]) Next() bool {
	if p.page >= p.lastPage {
		return false
	}
	*p = *NewLazy(p.loader, p.page+1, p.pageSize)
	return true
}

func (p *LazyPaginator[T]) ToJSON() ([]byte, error) {
	p.init()
	return p.Paginator.ToJSON()
}

func (p *LazyPaginator[T]) ToMap() map[string]any {
	p.init()
	return p.Paginator.ToMap()
}

func (p *LazyPaginator[T]) MarshalJSON() ([]byte, error) {
	p.init()
	return p.Paginator.MarshalJSON()
}

func (p *LazyPaginator[T]) UnmarshalJSON(_ []byte) error {
	return exception.NewUnsupportedException("json unmarshal is not supported")
}
