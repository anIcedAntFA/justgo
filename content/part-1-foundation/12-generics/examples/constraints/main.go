// Command constraints demonstrates the constraint ladder: any, comparable,
// cmp.Ordered, a custom type-union constraint, and the ~ approximation token
// that lets a constraint accept defined types built on a base type.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"cmp"
	"fmt"
)

// Print uses the loosest constraint, any (an alias for interface{}). With any,
// you can store, pass, and print T — but not use operators like + or <, because
// not every type supports them.
func Print[T any](value T) {
	fmt.Println(value)
}

// Contains needs == to compare elements, so T is constrained to comparable —
// the built-in constraint for types usable with == and != (and as map keys).
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// Sum needs +, so it uses a custom constraint. Number is an interface whose
// type set is a union of the numeric types. The ~ means "this type OR any type
// whose underlying type is this", so defined types like Celsius are accepted.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sum[T Number](nums []T) T {
	var total T // the zero value of T
	for _, n := range nums {
		total += n
	}
	return total
}

// Celsius is a DEFINED type with underlying type float64 (like Chapter 3's).
// Thanks to ~float64 in Number, Celsius satisfies the constraint.
type Celsius float64

func main() {
	Print(42)          // any: prints an int
	Print("hello")     // any: prints a string
	Print([]int{1, 2}) // any: prints a slice

	fmt.Println(Contains([]int{1, 2, 3}, 2))       // true  (comparable)
	fmt.Println(Contains([]string{"a", "b"}, "c")) // false

	fmt.Println(Sum([]int{1, 2, 3}))      // 6     (cmp of int)
	fmt.Println(Sum([]float64{1.5, 2.5})) // 4
	fmt.Println(Sum([]Celsius{20, 22}))   // 42    (~float64 in action)

	// cmp.Ordered (Go 1.21+) is the stdlib constraint for <, <=, >, >=.
	fmt.Println(cmp.Compare(3, 5)) // -1
}
