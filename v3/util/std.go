package util

// Any returns whether-or-not there is at least one item
// within the given slice that satisfies the given predicate.
//
// An empty slice always returns `false`.
func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

// All returns whether-or-not all items within the given slice
// satisfy the given predicate.
//
// An empty slice always returns `true`. This may seem counterintuitive,
// however it is due to being what is called a "vacuous truth". For
// more information, please see https://en.wikipedia.org/wiki/Vacuous_truth.
func All[T any](slice []T, predicate func(T) bool) bool {
	return !Any(slice, func(item T) bool {
		return !predicate(item)
	})
}
