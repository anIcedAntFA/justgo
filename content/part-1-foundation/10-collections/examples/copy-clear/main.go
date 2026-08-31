// Command copy-clear covers the tools for duplicating and emptying collections:
// the copy builtin, slices.Clone, and the clear builtin (Go 1.21) for both slices
// and maps.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"slices"
)

func main() {
	// copy(dst, src) copies min(len(dst), len(src)) elements and returns the count.
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3) // shorter on purpose
	n := copy(dst, src)
	fmt.Printf("copy → n=%d dst=%v\n", n, dst) // n=3 dst=[1 2 3]

	// slices.Clone (Go 1.21) is the concise, idiomatic full copy.
	clone := slices.Clone(src)
	clone[0] = 99
	fmt.Println("src stays:", src)   // [1 2 3 4 5]
	fmt.Println("clone:    ", clone) // [99 2 3 4 5]

	// clear on a slice zeroes every element (length is unchanged).
	clear(clone)
	fmt.Println("after clear(clone):", clone) // [0 0 0 0 0]

	// clear on a map deletes every key (length becomes 0).
	ages := map[string]int{"alice": 30, "bob": 25}
	clear(ages)
	fmt.Printf("after clear(ages): len=%d %v\n", len(ages), ages) // len=0 map[]
}
