// Command import-alias shows import aliasing — renaming a package at import time to
// resolve a name clash. Both math/rand/v2 and crypto/rand expose a package literally
// named "rand", so importing both unaliased would collide. An alias keeps them apart.
//
// Run it from this directory:
//
//	go run .
package main

import (
	crand "crypto/rand"
	"fmt"
	mrand "math/rand/v2"
)

func main() {
	// math/rand/v2 — fast pseudo-randomness for non-security uses like a dice roll.
	fmt.Println("dice roll:", mrand.IntN(6)+1)

	// crypto/rand — cryptographically secure randomness. Same short name "rand",
	// so the crand alias is what keeps the two imports from clashing.
	buf := make([]byte, 4)
	if _, err := crand.Read(buf); err != nil {
		fmt.Println("crypto/rand failed:", err)
		return
	}
	fmt.Printf("secure bytes: %x\n", buf)
}
