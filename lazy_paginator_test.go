package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLazy(t *testing.T) {
	loader := func(page, pageSize int) ([]string, int64) {
		items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		start := (page - 1) * pageSize
		if start >= len(items) {
			return []string{}, int64(len(items))
		}
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}
		return items[start:end], int64(len(items))
	}

	t.Run("valid page and page size", func(t *testing.T) {
		paginator := NewLazy(loader, 2, 3)
		assert.Equal(t, []string{"d", "e", "f"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 4, paginator.LastPage())
	})

	t.Run("zero page", func(t *testing.T) {
		paginator := NewLazy(loader, 0, 3)
		assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 4, paginator.LastPage())
	})

	t.Run("negative page", func(t *testing.T) {
		paginator := NewLazy(loader, -1, 3)
		assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 4, paginator.LastPage())
	})

	t.Run("zero page size", func(t *testing.T) {
		paginator := NewLazy(loader, 1, 0)
		assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 1, paginator.LastPage())
	})

	t.Run("negative page size", func(t *testing.T) {
		paginator := NewLazy(loader, 1, -3)
		assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 1, paginator.LastPage())
	})

	t.Run("page out of range", func(t *testing.T) {
		paginator := NewLazy(loader, 5, 3)
		assert.Empty(t, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 4, paginator.LastPage())
	})

	t.Run("empty data", func(t *testing.T) {
		emptyLoader := func(page, pageSize int) ([]string, int64) {
			return []string{}, 0
		}
		paginator := NewLazy(emptyLoader, 1, 10)
		assert.Empty(t, paginator.Items())
		assert.EqualValues(t, 0, paginator.Total())
		assert.Equal(t, 0, paginator.LastPage())
	})
}

func TestLazyPaginator_ToJSON(t *testing.T) {
	loader := func(page, pageSize int) ([]string, int64) {
		items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		start := (page - 1) * pageSize
		if start >= len(items) {
			return []string{}, int64(len(items))
		}
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}
		return items[start:end], int64(len(items))
	}

	paginator := NewLazy(loader, 2, 3)
	jsonBytes, err := paginator.ToJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, `{"items":["d","e","f"],"total":10,"lastPage":4,"page":2,"pageSize":3}`, string(jsonBytes))
}

func TestLazyPaginator_ToMap(t *testing.T) {
	loader := func(page, pageSize int) ([]string, int64) {
		items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		start := (page - 1) * pageSize
		if start >= len(items) {
			return []string{}, int64(len(items))
		}
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}
		return items[start:end], int64(len(items))
	}

	paginator := NewLazy(loader, 2, 3)
	assert.Equal(t, map[string]any{
		"items":    []string{"d", "e", "f"},
		"total":    int64(10),
		"lastPage": 4,
		"page":     2,
		"pageSize": 3,
	}, paginator.ToMap())
}

func TestLazyPaginator_MarshalJSON(t *testing.T) {
	loader := func(page, pageSize int) ([]string, int64) {
		items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		start := (page - 1) * pageSize
		if start >= len(items) {
			return []string{}, int64(len(items))
		}
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}
		return items[start:end], int64(len(items))
	}

	paginator := NewLazy(loader, 2, 3)
	jsonBytes, err := paginator.MarshalJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, `{"items":["d","e","f"],"total":10,"lastPage":4,"page":2,"pageSize":3}`, string(jsonBytes))
}

func TestLazyPaginator_UnmarshalJSON(t *testing.T) {
	loader := func(page, pageSize int) ([]string, int64) {
		items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		start := (page - 1) * pageSize
		if start >= len(items) {
			return []string{}, int64(len(items))
		}
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}
		return items[start:end], int64(len(items))
	}

	paginator := NewLazy(loader, 2, 3)
	err := paginator.UnmarshalJSON([]byte(`{"items":["d","e","f"],"total":10,"lastPage":4,"page":2,"pageSize":3}`))
	assert.EqualError(t, err, "Unsupported operation: json unmarshal is not supported")
}

func TestLazyPaginator_Next(t *testing.T) {
	loader := func(page, pageSize int) ([]string, int64) {
		items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		start := (page - 1) * pageSize
		if start >= len(items) {
			return []string{}, int64(len(items))
		}
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}
		return items[start:end], int64(len(items))
	}

	t.Run("move to next page", func(t *testing.T) {
		paginator := NewLazy(loader, 1, 3)
		assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
		assert.True(t, paginator.Next())
		assert.Equal(t, []string{"d", "e", "f"}, paginator.Items())
	})

	t.Run("no next page", func(t *testing.T) {
		paginator := NewLazy(loader, 4, 3)
		assert.Equal(t, []string{"j"}, paginator.Items())
		assert.False(t, paginator.Next())
		assert.Equal(t, []string{"j"}, paginator.Items())
	})

	t.Run("empty data", func(t *testing.T) {
		emptyLoader := func(page, pageSize int) ([]string, int64) {
			return []string{}, 0
		}
		paginator := NewLazy(emptyLoader, 1, 3)
		assert.Empty(t, paginator.Items())
		assert.False(t, paginator.Next())
		assert.Empty(t, paginator.Items())
	})
}
