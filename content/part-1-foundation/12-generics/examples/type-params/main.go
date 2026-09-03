// Command type-params demonstrates the core of Go generics: type parameters,
// type inference, and multiple type parameters. Compare each function to the
// duplicated / interface{} version it replaces.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"cmp"
	"fmt"
)

// Max is generic over T, constrained to cmp.Ordered (integers, floats, strings)
// so the > operator is available. One function replaces MaxInt/MaxFloat/MaxString.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Map transforms a []T into a []U using two type parameters. This is JS's
// Array.prototype.map, written once and fully type-safe.
func Map[T, U any](s []T, transform func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = transform(v)
	}
	return result
}

func main() {
	// Type inference: T is deduced from the arguments, so no [int] needed.
	fmt.Println(Max(3, 5))              // 5    → Max[int]
	fmt.Println(Max(3.2, 1.8))          // 3.2  → Max[float64]
	fmt.Println(Max("apple", "banana")) // banana → Max[string]
	fmt.Println(Max[int](3, 5))         // 5    → explicit, rarely needed

	// Map: T=int inferred from nums, U=string inferred from the func's return.
	nums := []int{1, 2, 3}
	labels := Map(nums, func(n int) string {
		return fmt.Sprintf("#%d", n)
	})
	fmt.Println(labels) // [#1 #2 #3]
}
