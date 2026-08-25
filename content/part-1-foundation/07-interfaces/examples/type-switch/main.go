// Command type-switch demonstrates the empty interface (any), the comma-ok type
// assertion, and the type switch — the three ways to recover a concrete type
// from an interface value.
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

// describe uses a type switch: inside each case, v already has that case's type.
func describe(i any) string {
	switch v := i.(type) {
	case nil:
		return "nil value"
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return fmt.Sprintf("string of length %d: %q", len(v), v)
	case bool:
		return fmt.Sprintf("bool: %t", v)
	case []int:
		return fmt.Sprintf("int slice with %d elements", len(v))
	default:
		return fmt.Sprintf("unknown type: %T", v)
	}
}

func main() {
	// Comma-ok assertion: safe, never panics. On failure the value is the
	// asserted type's zero value and ok is false.
	var i any = "hello"
	if s, ok := i.(string); ok {
		fmt.Printf("assertion ok: %q\n", s)
	}
	if n, ok := i.(int); !ok {
		fmt.Printf("not an int (n is the zero value: %d)\n", n)
	}

	fmt.Println("---")

	for _, v := range []any{42, "hi", true, []int{1, 2, 3}, 3.14, nil} {
		fmt.Println(describe(v))
	}
}
