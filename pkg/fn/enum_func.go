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
