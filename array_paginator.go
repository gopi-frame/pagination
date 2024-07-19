package pagination

import (
	"encoding/json"
)

type ArrayPaginator[T any] struct {
	Paginator[T]

	data []T
}

func NewArray[T any](data []T, page int, pageSize int) *ArrayPaginator[T] {
	page = max(1, page)
	if pageSize <= 0 {
		pageSize = 10
	}
	var paginator Paginator[T]
	total := len(data)
	if s := (page - 1) * pageSize; s >= total {
		paginator = *New[T]([]T{}, int64(total), page, pageSize)
	} else if e := page * pageSize; e >= total {
		paginator = *New[T](data[s:], int64(total), page, pageSize)
	} else {
		paginator = *New[T](data[s:e], int64(total), page, pageSize)
	}
	return &ArrayPaginator[T]{
		Paginator: paginator,
		data:      data,
	}
}

func (p *ArrayPaginator[T]) Next() bool {
	if p.Paginator.page >= p.Paginator.lastPage {
		return false
	}
	total := p.total
	page := p.Paginator.page + 1
	pageSize := p.pageSize
	data := p.data
	var paginator Paginator[T]
	if s := (page - 1) * pageSize; int64(s) >= total {
		paginator = *New[T]([]T{}, total, page, pageSize)
	} else if e := page * pageSize; int64(e) >= total {
		paginator = *New[T](data[s:], total, page, pageSize)
	} else {
		paginator = *New[T](data[s:e], total, page, pageSize)
	}
	p.Paginator = paginator
	return true
}

func (p *ArrayPaginator[T]) UnmarshalJSON(data []byte) error {
	var paginator Paginator[T]
	if err := json.Unmarshal(data, &paginator); err != nil {
		return err
	}
	p.Paginator = paginator
	return nil
}
