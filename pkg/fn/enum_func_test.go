package fn

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	sl := []int{1, 2, 3}
	result := Map(sl, func(i int) string { return strconv.Itoa(i) })
	assert.Equal(t, []int{1, 2, 3}, sl)
	assert.Equal(t, []string{"1", "2", "3"}, result)
}

func TestMapError(t *testing.T) {
	sl := []int{1, 2, 3}

	result, err := MapError(sl, func(i int) (string, error) {
		return strconv.Itoa(i), nil
	})

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, sl)
	assert.Equal(t, []string{"1", "2", "3"}, result)

	errExpected := errors.New("expected error")
	result, err = MapError(sl, func(i int) (string, error) {
		if i == 2 {
			return "", errExpected
		}
		return strconv.Itoa(i), nil
	})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, []int{1, 2, 3}, sl)
}

func TestFilter(t *testing.T) {
	sl := []int{1, 2, 3}
	result := Filter(sl, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2}, result)
	assert.Equal(t, []int{1, 2, 3}, sl)
}

func TestFilterError(t *testing.T) {
	sl := []int{1, 2, 3}
	result, err := FilterError(sl, func(i int) (bool, error) { return i%2 == 0, nil })
	assert.NoError(t, err)
	assert.Equal(t, []int{2}, result)
	assert.Equal(t, []int{1, 2, 3}, sl)

	errExpected := errors.New("expected error")
	result, err = FilterError(sl, func(i int) (bool, error) {
		if i == 2 {
			return false, errExpected
		}
		return i%2 == 1, nil
	})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, []int{1, 2, 3}, sl)
}

func TestReduce(t *testing.T) {
	sl := []int{1, 2, 3}
	result := Reduce(sl, func(a, b int) int {
		return a + b
	}, 0)
	assert.Equal(t, 6, result)
	assert.Equal(t, []int{1, 2, 3}, sl)

	result = Reduce(sl, func(a, b int) int {
		return a + b
	}, 2)
	assert.Equal(t, 8, result)
	assert.Equal(t, []int{1, 2, 3}, sl)
}
