package tools

func ZeroValue[T any]() T {
	var zero T
	return zero
}
