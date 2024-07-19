package pagination

import (
	"encoding/json"
	"math"
)

type Paginator[T any] struct {
	items    []T
	total    int64
	page     int
	pageSize int
	lastPage int
}

func New[T any](items []T, total int64, page int, pageSize int) *Paginator[T] {
	p := &Paginator[T]{
		items: items,
		total: total,
		page:  max(0, page),
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	p.pageSize = pageSize
	p.lastPage = int(math.Ceil(float64(total) / float64(pageSize)))
	return p
}

func (p *Paginator[T]) Items() []T {
	return p.items
}

func (p *Paginator[T]) Total() int64 {
	return p.total
}

func (p *Paginator[T]) LastPage() int {
	return p.lastPage
}

func (p *Paginator[T]) ToJSON() ([]byte, error) {
	return json.Marshal(p.ToMap())
}

func (p *Paginator[T]) ToMap() map[string]any {
	return map[string]any{
		"items":    p.items,
		"total":    p.total,
		"page":     p.page,
		"pageSize": p.pageSize,
		"lastPage": p.lastPage,
	}
}

func (p *Paginator[T]) MarshalJSON() ([]byte, error) {
	return p.ToJSON()
}

func (p *Paginator[T]) UnmarshalJSON(bytes []byte) error {
	type jsonObject struct {
		Items    []T   `json:"items"`
		Total    int64 `json:"total"`
		Page     int   `json:"page"`
		PageSize int   `json:"pageSize"`
		LastPage int   `json:"lastPage"`
	}
	var obj jsonObject
	if err := json.Unmarshal(bytes, &obj); err != nil {
		return err
	}
	p.items = obj.Items
	p.total = obj.Total
	p.page = obj.Page
	if obj.PageSize <= 0 {
		obj.PageSize = 10
	}
	p.pageSize = obj.PageSize
	p.lastPage = int(math.Ceil(float64(obj.Total) / float64(obj.PageSize)))
	return nil
}
