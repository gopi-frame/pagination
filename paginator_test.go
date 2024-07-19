package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	paginator := New[string]([]string{"a", "b", "c"}, 100, 1, 10)
	assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
	assert.EqualValues(t, 100, paginator.Total())
	assert.Equal(t, 10, paginator.LastPage())
}

func TestPaginator_Items_EmptySlice(t *testing.T) {
	paginator := New[string]([]string{}, 100, 1, 10)
	assert.Empty(t, paginator.Items())
}

func TestPaginator_Items_NilSlice(t *testing.T) {
	var slice []string
	paginator := New[string](slice, 100, 1, 10)
	assert.Nil(t, paginator.Items())
}

func TestPaginator_Items_InvalidPageSize(t *testing.T) {
	paginator := New[string]([]string{"a", "b", "c"}, 100, 1, 0)
	assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
}

func TestPaginator_Items_InvalidPageNumber(t *testing.T) {
	paginator := New[string]([]string{"a", "b", "c"}, 100, 0, 10)
	assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
}

func TestPaginator_ToMap(t *testing.T) {
	paginator := New[string]([]string{"a", "b", "c"}, 100, 1, 10)
	assert.Equal(t, map[string]interface{}{
		"items":    []string{"a", "b", "c"},
		"total":    int64(100),
		"lastPage": 10,
		"page":     1,
		"pageSize": 10,
	}, paginator.ToMap())
}

func TestPaginator_MarshalJSON(t *testing.T) {
	paginator := New[string]([]string{"a", "b", "c"}, 100, 1, 10)
	jsonBytes, err := paginator.MarshalJSON()
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	assert.JSONEq(t, `{"items":["a","b","c"],"total":100,"lastPage":10,"page":1,"pageSize":10}`, string(jsonBytes))
}

func TestPaginator_UnmarshalJSON(t *testing.T) {
	jsonBytes := []byte(`{"items":["a","b","c"],"total":100,"lastPage":10,"page":1,"pageSize":10}`)
	paginator := New[string]([]string{"a", "b", "c"}, 100, 1, 10)
	err := paginator.UnmarshalJSON(jsonBytes)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	assert.Equal(t, map[string]interface{}{
		"items":    []string{"a", "b", "c"},
		"total":    int64(100),
		"lastPage": 10,
		"page":     1,
		"pageSize": 10,
	}, paginator.ToMap())
}
