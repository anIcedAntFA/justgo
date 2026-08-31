// Command iterators shows Go 1.23 range-over-func: a function of the right shape
// (iter.Seq[T]) can be ranged over like a slice, enabling lazy custom sequences —
// conceptually similar to JavaScript generators.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

// Countdown returns an iterator that yields from..1. The yield callback returns
// false when the consumer breaks out of the loop, so the producer can stop early.
func Countdown(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	// Range over the function exactly like ranging over a collection.
	fmt.Print("countdown: ")
	for n := range Countdown(5) {
		fmt.Print(n, " ") // 5 4 3 2 1
	}
	fmt.Println()

	// Breaking early makes yield return false, and the producer stops.
	fmt.Print("first two: ")
	for n := range Countdown(5) {
		if n < 4 {
			break
		}
		fmt.Print(n, " ") // 5 4
	}
	fmt.Println()

	// The maps/slices packages expose iterators too. slices.Collect materializes
	// an iterator into a slice.
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := slices.Collect(maps.Keys(m))
	slices.Sort(keys)
	fmt.Println("keys:", keys) // [a b c]
}
