package fn

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var ErrEmptyValue = errors.New("empty value")

type Emptier interface {
	Empty() bool
}

type Option[T any] struct {
	value *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value: &value,
	}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsSome() bool {
	return o.value != nil
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) Value() (val T, err error) {
	if o.IsSome() {
		return *o.value, nil
	}
	return val, ErrEmptyValue
}

func (o Option[T]) ValueOrDefault(defaultValue T) T {
	if o.IsNone() {
		return defaultValue
	}
	return *o.value
}

func (o Option[T]) MustValue() T {
	if o.IsNone() {
		panic(ErrEmptyValue)
	}
	return *o.value
}

func (o Option[T]) ValueOrElse(fn func() T) T {
	if o.IsNone() {
		return fn()
	}
	return *o.value
}

func (o Option[T]) FlatMap(fn func(T) Option[T]) Option[T] {
	if o.IsNone() {
		return None[T]()
	}
	return fn(*o.value)
}

func (o Option[T]) Map(fn func(T) T) Option[T] {
	if o.IsNone() {
		return None[T]()
	}
	return Some(fn(*o.value))
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsNone() {
		var zero T
		return json.Marshal(zero)
	}
	return json.Marshal(o.value)
}
