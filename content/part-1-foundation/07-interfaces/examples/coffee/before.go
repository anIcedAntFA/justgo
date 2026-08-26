package main

import "fmt"

// BEFORE — no interface.
//
// Every brew method is a branch in one big switch, keyed on a string. It works,
// but look at what it costs you:
//
//  1. Open/Closed violation: adding a new brew method means editing THIS
//     function — reopening code that already worked and risking the branches
//     you did not touch. brewBAD grows without bound.
//  2. No compile-time safety: the method is a string. brewBAD("expreso", g) is
//     a typo the compiler happily accepts; it only misbehaves at run time (here,
//     it falls through to the default branch).
//  3. Not extensible from outside: a customer who brings their own brewing gear
//     cannot be served unless you come back and add a case for it.
func brewBAD(method string, g Grounds) Cup {
	switch method {
	case "phin":
		return Cup{Drink: "Cà phê phin", Origin: g.Origin, Strength: 8}
	case "espresso":
		return Cup{Drink: "Espresso", Origin: g.Origin, Strength: 9}
	case "coldbrew":
		return Cup{Drink: "Cold brew", Origin: g.Origin, Strength: 5}
	default:
		// A typo or an unknown method silently lands here.
		return Cup{Drink: fmt.Sprintf("??? (unknown method %q)", method), Origin: g.Origin}
	}
}
