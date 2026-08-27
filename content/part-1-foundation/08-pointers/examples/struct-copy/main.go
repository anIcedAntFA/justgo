// Command structcopy contrasts Go's value semantics with JavaScript's reference
// semantics: assigning a struct COPIES it, while assigning a pointer SHARES it.
package main

import "fmt"

// Counter is a plain struct — copied on assignment.
type Counter struct{ Count int }

func main() {
	// Struct assignment copies. Unlike a JS object, s2 is independent of s1.
	s1 := Counter{Count: 10}
	s2 := s1
	s2.Count = 20
	fmt.Println(s1.Count, s2.Count) // 10 20 — independent copies

	// Pointer assignment shares. p2 and p1 point at the same struct.
	p1 := &Counter{Count: 10}
	p2 := p1
	p2.Count = 20
	fmt.Println(p1.Count, p2.Count) // 20 20 — same underlying struct
}
