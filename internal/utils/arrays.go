package utils

func Map[T any, S any](source []T, fn func(T) S) []S {
	result := make([]S, len(source))

	for idx, item := range source {
		result[idx] = fn(item)
	}

	return result
}
