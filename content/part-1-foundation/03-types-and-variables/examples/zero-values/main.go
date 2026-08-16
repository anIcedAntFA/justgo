// Command zero-values demonstrates Go's zero values — the reason Go has no
// "undefined". Every variable is usable the moment it is declared.
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

func main() {
	var (
		i int
		f float64
		b bool
		s string
		p *int
	)

	fmt.Printf("int     → %d\n", i) // 0
	fmt.Printf("float64 → %g\n", f) // 0
	fmt.Printf("bool    → %t\n", b) // false
	fmt.Printf("string  → %q\n", s) // "" (empty, never nil)
	fmt.Printf("pointer → %v\n", p) // <nil>
}
