# Pagination

[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/pagination.svg)](https://pkg.go.dev/github.com/gopi-frame/pagination)
[![Go](https://github.com/gopi-frame/pagination/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/gopi-frame/pagination/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gopi-frame/pagination/graph/badge.svg?token=UGVGP6QF5O)](https://codecov.io/gh/gopi-frame/pagination)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/pagination)](https://goreportcard.com/report/github.com/gopi-frame/pagination)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
## Installation

```shell
go get -u -v github.com/gopi-frame/pagination
```

## Import
```go
import "github.com/gopi-frame/pagination"
```

## Usage

### Quick Start

```go
package main

import (
	"fmt"
	"github.com/gopi-frame/pagination"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	paginator := pagination.New(numbers, 10, 1, 5)
	fmt.Println(paginator.Items()) // [1, 2, 3, 4, 5]
	fmt.Println(paginator.LastPage()) // 2
	fmt.Println(paginator.Total()) // 10
}
```

### Array Pagination

`ArrayPaginator` is a simple pagination for array.
You can use method `Next()` to iterate through all pages.

```go
package main

import (
	"fmt"
	"github.com/gopi-frame/pagination"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pagination.NewArray[int](numbers, 1, 5)
	fmt.Println(paginator.Items()) // [1, 2, 3, 4, 5]
	fmt.Println(paginator.LastPage()) // 2
	fmt.Println(paginator.Total()) // 10
	// iterate
	for pagigator.Next() {
		// do something
    }
}
```

### Lazy Load Pagination

`LazyPaginator` is a pagination for lazy load data.
You can use method `Next()` to iterate through all pages.

```go
package main

import "github.com/gopi-frame/pagination"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var loader = func(page, pageSize int) ([]int, int64) {
		total := int64(len(numbers))
		s := (page - 1) * pageSize
		if int64(s) >= total {
			return []int{}, total
		}
		e := page * pageSize
		if e >= total {
			return numbers[s:], total
		}
		return numbers[s:e], total
	}
	paginator := pagination.NewLazy(loader, 10, 1)
	// iterate
	for pagigator.Next() {
		// do something
	}
}
```