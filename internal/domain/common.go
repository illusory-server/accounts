package domain

import (
	"github.com/pkg/errors"
)

var ErrEmptyOptionalValue = errors.New("empty optional value")

type Option[T any] struct {
	value *T
}

func (o Option[T]) Empty() bool {
	return o.value == nil
}

func (o Option[T]) ValueOrDefault(defaultValue T) T {
	if o.Empty() {
		return defaultValue
	}
	return *o.value
}

func (o Option[T]) Value() (val T, err error) {
	if o.Empty() {
		return val, ErrEmptyOptionalValue
	}
	return *o.value, nil
}

func (o Option[T]) Set(value T) Option[T] {
	o.value = &value
	return o
}

func NewOptional[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func NewEmptyOptional[T any]() Option[T] {
	return Option[T]{value: nil}
}
