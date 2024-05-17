package pagination

// ArrayPaginator is a paginator based on an array, it extends from [Paginator]
//
// Example:
//   allItems := []int{}
//   for i := 0; i < 100; i++ {
//       allItems = append(allItems, i)
//   }
//
//   pageSize := 10
//   page := 1
//
//   paginator := Array[int](allItems, pageSize, page)
type ArrayPaginator[T any] struct {
	*Paginator[T]

	all []T
}

// NewArrayPaginator creates an [ArrayPaginator] instance
func NewArrayPaginator[T any](all []T, pageSize, page int) *ArrayPaginator[T] {
	paginator := new(ArrayPaginator[T])
	paginator.Paginator = new(Paginator[T])
	paginator.all = all
	if page < 0 {
		page = 1
	}
	if pageSize < 0 {
		pageSize = 10
	}
	paginator.total = int64(len(all))
	paginator.currentPage = page
	paginator.pageSize = pageSize
	var values []T
	startIndex := int64((page - 1) * pageSize)
	endIndex := startIndex + int64(pageSize)
	if startIndex > paginator.total {
		values = make([]T, 0)
	} else if endIndex > paginator.total {
		values = all[startIndex:]
	} else {
		values = all[startIndex:endIndex]
	}
	paginator.Paginator = NewPaginator(int64(len(all)), values, pageSize, page)
	return paginator
}

// All returns all items
func (p *ArrayPaginator[T]) All() []T {
	return p.all
}
