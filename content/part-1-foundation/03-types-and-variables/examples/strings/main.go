// Command strings is the Chapter 03 "String Deep Dive" example: a Go string is
// UTF-8 encoded bytes, so len() counts bytes while a range loop yields runes
// (Unicode code points). Note the emoji 🌍 takes 4 bytes.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界! 🌍"

	fmt.Printf("string              → %q\n", s)
	fmt.Printf("length in bytes     → %d\n", len(s))                    // 19
	fmt.Printf("length in runes     → %d\n", utf8.RuneCountInString(s)) // 12

	fmt.Println("\nbyte index | rune | code point")
	for i, r := range s {
		fmt.Printf("%10d | %-4c | U+%04X\n", i, r, r)
	}
}
