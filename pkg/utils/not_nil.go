package utils

func NotNil[T comparable](v T) T {
	var t T
	if v == t {
		panic("nil value is not valid")
	}
	return v
}
