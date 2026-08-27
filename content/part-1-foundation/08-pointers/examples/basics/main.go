// Command basics shows the two pointer operators (& and *), automatic
// dereference on struct fields, and that a pointer's zero value is nil.
package main

import "fmt"

// User is a small struct we take the address of below.
type User struct {
	Name string
	Age  int
}

func main() {
	x := 42
	p := &x // & — "address of": p is a *int pointing at x

	fmt.Println(x)  // 42  — the value
	fmt.Println(*p) // 42  — * — "value at": dereference p

	*p = 100
	fmt.Println(x) // 100 — x changed because p points at it

	// Automatic dereference: through a *User you write u.Name, never (*u).Name.
	u := &User{Name: "Alice", Age: 30}
	fmt.Println(u.Name) // Alice
	u.Age = 31          // modifies the struct through the pointer
	fmt.Println(u.Age)  // 31

	// The zero value of any pointer is nil — it points at nothing.
	var np *int
	fmt.Println(np == nil) // true
}
