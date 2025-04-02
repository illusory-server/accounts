package tools

// ZeroValue all generic type zero value
func ZeroValue[T any]() T {
	var zero T
	return zero
}
