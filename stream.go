package main

func Map[T any, B any](s []T, mapFunc func(T) B) []B {
	mapped := make([]B, 0, len(s))
	for _, t := range s {
		mapped = append(mapped, mapFunc(t))
	}
	return mapped
}

func Filter[T any](s []T, filterFunc func(T) bool) []T {
	filtered := make([]T, 0, len(s))
	for _, t := range s {
		if filterFunc(t) {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

// can not return a nil or any, so it is not possible to go from T slice to singe T when the value could possible not be found
/*func FindFirst[T any](s []T, filterFunc func(T) bool) T {
	for _, t := range s {
		if filterFunc(t) {
			return t
		}
	}
	return nil
}
*/
