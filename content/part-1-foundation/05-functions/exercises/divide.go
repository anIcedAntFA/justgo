// Package exercises holds the coding exercises for Chapter 05: Functions.
//
// How to use: read the TODO, implement the function, then remove the t.Skip in
// the matching _test.go and run `go test ./...` until it passes.
package exercises

import "errors"

// Divide returns the quotient a/b. On division by zero it must return a zero
// quotient and a non-nil error; otherwise the quotient and a nil error. This is
// the (value, error) backbone in miniature.
//
// TODO: implement this. Use errors.New (or fmt.Errorf) for the zero case and add
// the import yourself.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}
