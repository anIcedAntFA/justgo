// Command mutation shows why Go needs pointers to modify a caller's value:
// arguments are passed by value (copied), so only a pointer reaches the original.
package main

import "fmt"

// incValue receives a COPY of n. It increments and prints the copy to show the
// change is real — but purely local; the caller's variable never sees it.
func incValue(n int) {
	n++
	fmt.Println("inside incValue, local n =", n) // 11 — the copy
}

// incPointer receives the ADDRESS of n — it modifies the original.
func incPointer(n *int) { *n++ }

func main() {
	x := 10

	incValue(x)
	fmt.Println(x) // 10 — the copy was modified, original untouched

	incPointer(&x)
	fmt.Println(x) // 11 — original modified through the pointer
}
