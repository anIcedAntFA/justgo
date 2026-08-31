// Command slice-header makes the length/capacity distinction visible and shows
// how append grows the backing array. A slice is a small 3-word header (pointer,
// len, cap); the data lives in a separate backing array.
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

func main() {
	// make([]T, len, cap): length 3 you can read now, capacity 10 to grow into.
	s := make([]int, 3, 10)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=3 cap=10 [0 0 0]

	// Appending within capacity reuses the same backing array — no reallocation.
	s = append(s, 1)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=4 cap=10

	// Watch capacity grow as we append past it. Go grows amortized (roughly
	// doubling while small, then more gently) and rounds up to allocator size
	// classes — so the exact cap steps are runtime-defined, not a promise.
	fmt.Println("\ngrowth from an empty slice:")
	var g []int
	prev := cap(g)
	for i := range 12 { // Go 1.22+: range over an integer
		g = append(g, i)
		if cap(g) != prev {
			fmt.Printf("  len=%2d cap=%2d  <- backing array regrew\n", len(g), cap(g))
			prev = cap(g)
		}
	}

	// Preallocating when you know the size avoids every regrowth above.
	fmt.Println("\npreallocated with make([]int, 0, 12):")
	p := make([]int, 0, 12)
	prev = cap(p)
	for i := range 12 {
		p = append(p, i)
		if cap(p) != prev {
			fmt.Printf("  len=%2d cap=%2d\n", len(p), cap(p))
			prev = cap(p)
		}
	}
	fmt.Println("  (never regrew — one allocation up front)")
}
