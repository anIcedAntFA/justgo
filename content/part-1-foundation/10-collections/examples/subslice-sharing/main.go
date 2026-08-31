// Command subslice-sharing demonstrates Go's nastiest slice gotcha: a subslice
// shares the parent's backing array, so writing through one is visible through
// the other — and appending into spare capacity silently corrupts the parent.
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
	// 1. Slicing gives a WINDOW into the same memory, not a copy.
	original := []int{1, 2, 3, 4, 5}
	sub := original[1:4] // [2 3 4]
	sub[0] = 99          // writes into original[1]
	fmt.Println("after sub[0]=99:")
	fmt.Println("  original:", original) // [1 99 3 4 5] — changed!
	fmt.Println("  sub:     ", sub)      // [99 3 4]

	// 2. The append trap: a subslice keeps the parent's spare capacity, so
	// appending writes IN PLACE, over the parent's data.
	parent := []int{1, 2, 3, 4, 5}
	head := parent[0:2] // len 2, but cap 5 (inherits room to the right)
	fmt.Printf("\nhead=%v len=%d cap=%d\n", head, len(head), cap(head))
	head = append(head, 999)         // there's spare cap → overwrites parent[2]
	fmt.Println("  head:  ", head)   // [1 2 999]
	fmt.Println("  parent:", parent) // [1 2 999 4 5] — corrupted!

	// 3a. Fix with slices.Clone — an independent backing array.
	safeParent := []int{1, 2, 3, 4, 5}
	cloned := slices.Clone(safeParent[0:2])
	cloned = append(cloned, 999)
	fmt.Printf("\nclone fix     — parent: %v  cloned: %v\n", safeParent, cloned) // parent [1 2 3 4 5] safe

	// 3b. Fix with the three-index full slice expression s[low:high:max], which
	// caps the result so the next append is forced to reallocate.
	capped := safeParent[0:2:2] // cap == 2 now
	capped = append(capped, 999)
	fmt.Printf("three-index fix — parent: %v  capped: %v\n", safeParent, capped) // parent [1 2 3 4 5] safe
}
