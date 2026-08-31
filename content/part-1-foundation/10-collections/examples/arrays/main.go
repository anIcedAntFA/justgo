// Command arrays shows why arrays are rarely used directly: their length is part
// of their type, and they are VALUES — assigning or passing one copies it whole.
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

func main() {
	// The length is part of the type: [3]int and [4]int are different types.
	var a [3]int
	a[0], a[1] = 10, 20
	fmt.Println("a:", a, "len:", len(a)) // a: [10 20 0] len: 3

	// Let the compiler count the elements with [...].
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("c = %v is a [%d]int\n", c, len(c)) // c = [1 2 3 4 5] is a [5]int

	// Arrays are values: assigning COPIES every element.
	b := a
	b[0] = 99
	fmt.Println("a:", a) // a: [10 20 0] — unchanged
	fmt.Println("b:", b) // b: [99 20 0]

	// Passing to a function copies too — the callee can't mutate the caller's array.
	zero(a)
	fmt.Println("a after zero(a):", a) // still [10 20 0]
}

// zero sets every element to 0 — but on its own COPY of the array.
func zero(arr [3]int) {
	for i := range arr {
		arr[i] = 0
	}
}
