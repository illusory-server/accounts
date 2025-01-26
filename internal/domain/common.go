package domain

type Option[T any] struct {
	value *T
}

func (o Option[T]) Empty() bool {
	return o.value == nil
}

func (o Option[T]) ValueOrDefault(defaultValue T) T {
	if o.value == nil {
		return defaultValue
	}
	return *o.value
}

func (o Option[T]) Set(value T) {
	o.value = &value
}

func NewOptional[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func NewEmptyOptional[T any]() Option[T] {
	return Option[T]{value: nil}
}
