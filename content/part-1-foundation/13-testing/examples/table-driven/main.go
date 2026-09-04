// Command table-driven shows the canonical Go test pattern. The runnable main
// is just a demo of the function under test; the real lesson is in
// main_test.go — a table of cases looped with t.Run subtests, including an
// error case.
//
// Run the demo:
//
//	go run .
//
// Run the tests (the point of the example):
//
//	go test -v
package main

import (
	"errors"
	"fmt"
)

// ErrDivideByZero is returned by Divide when the divisor is zero. Testing a
// specific sentinel error (Chapter 9) with errors.Is is cleaner than a bare
// wantErr bool.
var ErrDivideByZero = errors.New("divide by zero")

// Classify buckets a number into a short label. Several branches means several
// edge cases — exactly what a table test is good at.
func Classify(n int) string {
	switch {
	case n == 0:
		return "zero"
	case n < 0:
		return "negative"
	case n%2 == 0:
		return "positive-even"
	default:
		return "positive-odd"
	}
}

// Divide returns a/b, or ErrDivideByZero when b is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	for _, n := range []int{0, -4, 6, 7} {
		fmt.Printf("Classify(%d) = %s\n", n, Classify(n))
	}

	q, err := Divide(10, 2)
	fmt.Printf("Divide(10, 2) = %v, err = %v\n", q, err)

	_, err = Divide(1, 0)
	fmt.Printf("Divide(1, 0) err = %v\n", err)
}
