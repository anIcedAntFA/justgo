// Command nil shows the most common Go runtime panic — dereferencing a nil
// pointer — and the explicit nil guard that prevents it. Go has no optional
// chaining (?.), so you check nil yourself.
package main

import "fmt"

// User is the value a *User might or might not point at.
type User struct{ Name string }

// nameOf guards against nil before touching any field.
func nameOf(u *User) string {
	if u == nil {
		return "<no user>" // guard first — u.Name below would panic on nil
	}
	return u.Name
}

func main() {
	var u *User            // nil — points at nothing
	fmt.Println(nameOf(u)) // <no user>

	u = &User{Name: "Alice"}
	fmt.Println(nameOf(u)) // Alice

	// Dereferencing nil panics. We recover here only to show it without
	// crashing the program — real code prevents it with the guard above.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered:", r)
			}
		}()
		var bad *User
		fmt.Println(bad.Name) // panic: invalid memory address or nil pointer dereference
	}()
}
