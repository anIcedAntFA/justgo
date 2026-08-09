// Command hello is the Chapter 02 "first program" — it verifies your Go
// toolchain end to end: build, run, format, and test.
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

func main() {
	fmt.Println("Hello from Go!")

	// Multiple return values — you'll see this everywhere in Go.
	name, age := "Go Developer", 2026-2007
	fmt.Printf("I'm a %s, Go is %d years old\n", name, age)
}
