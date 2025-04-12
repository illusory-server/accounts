package fn

func Map[T, K any](slice []T, fn func(T) K) []K {
	result := make([]K, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func MapError[T, K any](slice []T, fn func(T) (K, error)) ([]K, error) {
	result := make([]K, len(slice))
	for i, v := range slice {
		k, err := fn(v)
		if err != nil {
			return nil, err
		}
		result[i] = k
	}
	return result, nil
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterError[T any](slice []T, fn func(T) (bool, error)) ([]T, error) {
	result := make([]T, 0)
	for _, v := range slice {
		ok, err := fn(v)
		if err != nil {
			return nil, err
		}
		if ok {
			result = append(result, v)
		}
	}
	return result, nil
}

func Reduce[T any, K any](slice []T, fn func(acc K, item T) K, initial K) K {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}
