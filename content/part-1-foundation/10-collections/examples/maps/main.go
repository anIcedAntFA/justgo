// Command maps covers the map essentials that trip up JS developers: the zero
// value on a missing key, the comma-ok existence check, the nil-map write panic,
// randomized iteration order, and the map-of-pointers workaround for mutating
// struct values in place.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"maps"
	"slices"
)

type point struct{ X, Y int }

func main() {
	ages := map[string]int{"alice": 30, "bob": 25}

	// A missing key returns the ZERO VALUE, not "undefined" and not an error.
	fmt.Println("zoe:", ages["zoe"]) // 0

	// comma-ok tells "present with value 0" apart from "absent".
	if age, ok := ages["alice"]; ok {
		fmt.Println("alice is", age) // alice is 30
	}
	_, ok := ages["zoe"]
	fmt.Println("zoe present?", ok) // false

	// delete is a no-op if the key is absent.
	delete(ages, "bob")

	// Deterministic output needs sorted keys — map iteration order is randomized.
	// maps.Keys (Go 1.23) yields an iterator; slices.Sorted collects+sorts it.
	for _, k := range slices.Sorted(maps.Keys(ages)) {
		fmt.Printf("  %s: %d\n", k, ages[k])
	}

	// A nil map reads fine but PANICS on write — always make (or use a literal).
	var counts map[string]int
	fmt.Println("nil-map read:", counts["x"]) // 0, no panic
	counts = make(map[string]int)
	counts["x"]++ // now safe

	// Map values are not addressable, so you can't assign to a struct field in a
	// value map: m["a"].X = 1 would not compile. Use a map of POINTERS instead.
	shapes := map[string]*point{"a": {1, 2}}
	shapes["a"].X = 10                      // works — modifying through the pointer
	fmt.Println("shapes[a]:", *shapes["a"]) // {10 2}
}
