// Command escape shows that returning a pointer to a local variable is safe in
// Go. The compiler's escape analysis sees the address outlives the function and
// allocates the variable on the heap. Observe the decision with:
//
//	go build -gcflags=-m ./content/part-1-foundation/08-pointers/examples/escape
//
// (look for "moved to heap: u" / "&u escapes to heap").
package main

import "fmt"

// User is allocated inside newUser and outlives the call.
type User struct{ Name string }

// newUser returns the address of a LOCAL variable. In C this dangles; in Go
// escape analysis moves u to the heap, so the pointer stays valid.
func newUser(name string) *User {
	u := User{Name: name} // local
	return &u             // escapes — safe, GC frees it later
}

func main() {
	a := newUser("Alice")
	b := newUser("Bob")

	fmt.Println(a.Name, b.Name) // Alice Bob
	fmt.Println(a != b)         // true — two distinct allocations
}
