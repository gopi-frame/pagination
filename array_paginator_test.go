package pagination

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArray(t *testing.T) {
	data := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	t.Run("valid page and page size", func(t *testing.T) {
		paginator := NewArray(data, 2, 3)
		assert.Equal(t, []string{"d", "e", "f"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 4, paginator.LastPage())
	})

	t.Run("zero page", func(t *testing.T) {
		paginator := NewArray(data, 0, 3)
		assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 4, paginator.LastPage())
	})

	t.Run("negative page", func(t *testing.T) {
		paginator := NewArray(data, -1, 3)
		assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 4, paginator.LastPage())
	})

	t.Run("zero page size", func(t *testing.T) {
		paginator := NewArray(data, 1, 0)
		assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 1, paginator.LastPage())
	})

	t.Run("negative page size", func(t *testing.T) {
		paginator := NewArray(data, 1, -3)
		assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 1, paginator.LastPage())
	})

	t.Run("page out of range", func(t *testing.T) {
		paginator := NewArray(data, 5, 3)
		assert.Empty(t, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 4, paginator.LastPage())
	})

	t.Run("empty data", func(t *testing.T) {
		paginator := NewArray([]string{}, 1, 10)
		assert.Empty(t, paginator.Items())
		assert.EqualValues(t, 0, paginator.Total())
		assert.Equal(t, 0, paginator.LastPage())
	})
}

func TestArrayPaginator_UnmarshalJSON(t *testing.T) {
	t.Run("valid JSON", func(t *testing.T) {
		data := []byte(`{"items":["a","b","c"],"total":10,"lastPage":4,"page":2,"pageSize":3}`)
		var paginator ArrayPaginator[string]
		err := json.Unmarshal(data, &paginator)
		assert.NoError(t, err)
		assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
		assert.EqualValues(t, 10, paginator.Total())
		assert.Equal(t, 4, paginator.LastPage())
	})

	t.Run("invalid JSON", func(t *testing.T) {
		data := []byte(`{"items":["a","b","c"],"total":"invalid","last_page":4,"page":2,"page_size":3}`)
		var paginator ArrayPaginator[string]
		err := json.Unmarshal(data, &paginator)
		assert.Error(t, err)
	})

	t.Run("empty JSON", func(t *testing.T) {
		data := []byte(`{}`)
		var paginator ArrayPaginator[string]
		err := json.Unmarshal(data, &paginator)
		assert.NoError(t, err)
		assert.Empty(t, paginator.Items())
		assert.EqualValues(t, 0, paginator.Total())
		assert.Equal(t, 0, paginator.LastPage())
	})

}
func TestArrayPaginator_Next(t *testing.T) {
	data := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	t.Run("move to next page", func(t *testing.T) {
		paginator := NewArray(data, 1, 3)
		assert.Equal(t, []string{"a", "b", "c"}, paginator.Items())
		assert.True(t, paginator.Next())
		assert.Equal(t, []string{"d", "e", "f"}, paginator.Items())
	})

	t.Run("no next page", func(t *testing.T) {
		paginator := NewArray(data, 4, 3)
		assert.Equal(t, []string{"j"}, paginator.Items())
		assert.False(t, paginator.Next())
		assert.Equal(t, []string{"j"}, paginator.Items())
	})

	t.Run("empty data", func(t *testing.T) {
		paginator := NewArray([]string{}, 1, 3)
		assert.Empty(t, paginator.Items())
		assert.False(t, paginator.Next())
		assert.Empty(t, paginator.Items())
	})
}
