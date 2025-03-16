package fn

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type inner struct {
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

type myType struct {
	Field Option[inner] `json:"field"`
	Name  string        `json:"name"`
}

func TestOptional(t *testing.T) {
	t.Run("correct constructor", func(t *testing.T) {
		opt := Some(5)
		expected := 5
		assert.Equal(t, Option[int]{
			value: &expected,
		}, opt)

		opt = None[int]()
		assert.Equal(t, Option[int]{}, opt)
	})

	t.Run("Should correct check method", func(t *testing.T) {
		opt := Some(5)
		assert.True(t, opt.IsSome())
		assert.False(t, opt.IsNone())

		opt = None[int]()
		assert.True(t, opt.IsNone())
		assert.False(t, opt.IsSome())
	})

	t.Run("Should correct get value", func(t *testing.T) {
		opt := Some(5)
		res, err := opt.Value()
		assert.NoError(t, err)
		assert.Equal(t, 5, res)

		res = opt.ValueOrDefault(2)
		assert.Equal(t, 5, res)

		res = opt.ValueOrElse(func() int {
			return 1
		})
		assert.Equal(t, 5, res)

		opt = None[int]()
		res, err = opt.Value()
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrEmptyValue))
		assert.Equal(t, 0, res)

		res = opt.ValueOrDefault(2)
		assert.Equal(t, 2, res)
		res = opt.ValueOrElse(func() int {
			return 1
		})
		assert.Equal(t, 1, res)
	})

	t.Run("Should correct must get value", func(t *testing.T) {
		opt := Some(5)
		res := opt.MustValue()
		assert.Equal(t, 5, res)

		defer func() {
			r := recover()
			assert.NotNil(t, r)
			e, ok := r.(error)
			assert.True(t, ok)
			assert.Equal(t, ErrEmptyValue.Error(), e.Error())
		}()
		opt = None[int]()
		opt.MustValue()
	})

	t.Run("Should correct map", func(t *testing.T) {
		opt := Some(5)
		res, err := opt.Value()
		assert.NoError(t, err)
		assert.Equal(t, 5, res)
		opt = opt.FlatMap(func(i int) Option[int] {
			return Some(i + 4)
		})
		res, err = opt.Value()
		assert.NoError(t, err)
		assert.Equal(t, 5+4, res)

		opt = opt.Map(func(i int) int {
			return i + 6
		})
		res, err = opt.Value()
		assert.NoError(t, err)
		assert.Equal(t, 15, res)

		opt = None[int]()
		opt = opt.FlatMap(func(i int) Option[int] {
			return Some(i + 4)
		})
		assert.Equal(t, None[int](), opt)

		opt = opt.Map(func(i int) int {
			return i + 4
		})
		assert.Equal(t, None[int](), opt)
	})
}

func TestOptionalJson(t *testing.T) {
	m := myType{
		Name: "keka",
		Field: Some(inner{
			Value1: 1,
			Value2: 2,
		}),
	}
	expectedJson := `{"field":{"value1":1,"value2":2},"name":"keka"}`
	data, err := json.Marshal(m)
	assert.NoError(t, err)
	assert.Equal(t, expectedJson, string(data))

	m = myType{
		Name:  "keka",
		Field: None[inner](),
	}
	expectedJson = `{"field":{"value1":0,"value2":0},"name":"keka"}`
	data, err = json.Marshal(m)
	assert.NoError(t, err)
	assert.Equal(t, expectedJson, string(data))
}
