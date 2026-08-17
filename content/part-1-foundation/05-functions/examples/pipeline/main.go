// Command pipeline is the Chapter 05 capstone: it threads together first-class
// functions passed as arguments, a variadic transformer, a closure that carries
// state, and a range-over-func iterator (iter.Seq, Go 1.23+) — one small
// program touching every idea in the chapter.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"iter"
	"strings"
)

// transform applies a chain of first-class functions left-to-right to n.
// fns is variadic, so callers pass as many stages as they like.
func transform(n int, fns ...func(int) int) int {
	for _, f := range fns {
		n = f(n)
	}
	return n
}

// tagger returns a closure that prefixes each label with an incrementing number,
// carrying its own count across calls.
func tagger(prefix string) func(string) string {
	count := 0
	return func(label string) string {
		count++
		return fmt.Sprintf("%s%d:%s", prefix, count, label)
	}
}

// countTo returns an iterator (iter.Seq[int]) — an ordinary function whose one
// parameter is a yield callback. `for range` drives it. yield reports false when
// the loop breaks early, which lets the iterator stop and clean up.
func countTo(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	double := func(n int) int { return n * 2 }
	addOne := func(n int) int { return n + 1 }

	// First-class functions + variadic: build a pipeline from small stages.
	fmt.Println("transform(5, double, addOne):", transform(5, double, addOne)) // (5*2)+1 = 11

	// Closure carrying state across calls.
	tag := tagger("row-")
	fmt.Println(tag("alpha")) // row-1:alpha
	fmt.Println(tag("beta"))  // row-2:beta

	// Range over a function iterator (Go 1.23+), feeding each value through the pipeline.
	var b strings.Builder
	for v := range countTo(3) {
		fmt.Fprintf(&b, "%d ", transform(v, double, addOne))
	}
	fmt.Println("piped iterator:", strings.TrimSpace(b.String())) // 3 5 7
}
