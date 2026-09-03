// Package exercises holds the coding exercises for Chapter 12: Generics.
//
// How to use: read the TODO, implement the function or type, then remove the
// t.Skip in the matching _test.go and run `go test ./...` until it passes.
package exercises

// Map returns a new slice holding fn(v) for each v in s. It is generic over two
// type parameters: T (the input element) and U (the output element).
//
// TODO: allocate a []U of len(s) and fill it by applying fn to each element.
func Map[T, U any](s []T, fn func(T) U) []U {
	return nil // TODO
}

// Filter returns a new slice holding only the elements of s for which keep
// returns true. One type parameter is enough — input and output share type T.
//
// TODO: append each element that keep(v) accepts to a fresh result slice.
func Filter[T any](s []T, keep func(T) bool) []T {
	return nil // TODO
}

// Reduce folds s into a single U value, starting from initial and combining
// each element with the accumulator via fn.
//
// TODO: start acc at initial, walk s applying acc = fn(acc, v), return acc.
func Reduce[T, U any](s []T, initial U, fn func(U, T) U) U {
	return initial // TODO
}
