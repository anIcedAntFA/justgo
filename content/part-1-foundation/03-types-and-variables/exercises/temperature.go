// Package exercises holds the coding exercises for Chapter 03: Types & Variables.
//
// How to use: read the TODO, implement the function, then remove the t.Skip in
// the matching _test.go and run `go test ./...` until it passes.
package exercises

// Celsius and Fahrenheit are DEFINED types (not aliases) on top of float64.
// Being distinct types, the compiler refuses to mix them: you cannot pass a
// Celsius where a Fahrenheit is wanted, nor add the two, without an explicit
// conversion. That compile-time safety for domain values is the whole point.
//
//	var temp Celsius = Fahrenheit(100) // ❌ does not compile — different types
type (
	Celsius    float64
	Fahrenheit float64
)

// CtoF converts a temperature from Celsius to Fahrenheit.
//
// Formula: F = C*9/5 + 32
//
// The arithmetic stays in Celsius (a float64) until the final conversion, so
// 9/5 is float division, not integer division (which would be 1).
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FtoC converts a temperature from Fahrenheit to Celsius.
//
// Formula: C = (F - 32) * 5/9
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
