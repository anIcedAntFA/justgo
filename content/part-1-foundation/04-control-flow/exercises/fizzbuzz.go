// Package exercises holds the coding exercises for Chapter 04: Control Flow.
//
// How to use: read the TODO, implement the function, then remove the t.Skip in
// the matching _test.go and run `go test ./...` until it passes.
package exercises

import "strconv"

// FizzBuzz returns the FizzBuzz word for a single number n:
//
//	divisible by 3 and 5 → "FizzBuzz"
//	divisible by 3       → "Fizz"
//	divisible by 5       → "Buzz"
//	otherwise            → the number as a string (e.g. 7 → "7")
//
// TODO: implement this with a TAGLESS switch (no if/else). For the default case
// you'll want strconv.Itoa(n) — add the "strconv" import yourself.
func FizzBuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}
