package contract

// Paginator is the paginator interface
type Paginator[T any] interface {
	// Items returns all items of current page
	Items() []T
	// Total returns the total count
	Total() int64
	// FirstItemIndex returns the index of the first item of current page
	FirstItemIndex() int
	// LastItemIndex returns the index of the last item of current page
	LastItemIndex() int
	// CurrentPage returns the current page number
	CurrentPage() int
	// PageSize returns the page size
	PageSize() int
	// HasMore returns whether the next page exists
	HasMore() bool
	// LastPage returns the last page number
	LastPage() int
	// ToMap returns current paginator info as a map
	ToMap() map[string]any
}
