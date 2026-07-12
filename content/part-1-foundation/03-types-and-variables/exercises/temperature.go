// Package exercises holds the coding exercises for Chapter 03: Types & Variables.
//
// How to use: read the TODO, implement the function, then remove the t.Skip in
// the matching _test.go and run `go test ./...` until it passes.
package exercises

// CelsiusToFahrenheit converts a temperature in Celsius to Fahrenheit.
//
// Formula: F = C*9/5 + 32
//
// TODO: implement this. Watch the types — Go does NO implicit numeric
// conversion, and 9/5 with integer literals is integer division (== 1). Keep the
// arithmetic in float64.
func CelsiusToFahrenheit(c float64) float64 {
	return c // TODO: replace with the real formula
}
