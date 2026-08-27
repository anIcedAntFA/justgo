// Command newvslit compares the two ways to get a pointer to fresh storage:
// the built-in new(T) and the composite literal &T{}. Reach for &T{} by default
// — it lets you initialize fields; new is handy for a zeroed primitive.
package main

import "fmt"

// User has fields we want to initialize inline below.
type User struct {
	Name string
	Age  int
}

func main() {
	// new(T): allocate a zeroed T, return *T. Cannot set fields inline.
	p := new(int)
	fmt.Println(*p) // 0
	*p = 42
	fmt.Println(*p) // 42

	u1 := new(User)          // zeroed
	fmt.Printf("%+v\n", *u1) // {Name: Age:0}

	// &T{}: the idiomatic form — initialize fields as you allocate.
	u2 := &User{Name: "Alice", Age: 30}
	fmt.Printf("%+v\n", *u2) // {Name:Alice Age:30}
}
