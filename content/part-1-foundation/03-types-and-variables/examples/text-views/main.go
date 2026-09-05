// Command text-views is the Chapter 03 "Strings, Bytes & Runes" deep-dive example:
// the same text seen three ways (string / []byte / []rune), why len differs, how
// range decodes UTF-8, and building output efficiently with strings.Builder.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "café"

	// Three views of the same text. Converting copies (strings are immutable).
	b := []byte(s)
	r := []rune(s)

	fmt.Printf("string %q\n", s)
	fmt.Printf("len(string) = %d bytes\n", len(s)) // 5 — é is 2 bytes
	fmt.Printf("len([]byte) = %d bytes\n", len(b)) // 5
	fmt.Printf("len([]rune) = %d runes\n", len(r)) // 4 — c a f é

	// Indexing a string yields a BYTE (uint8), not a character.
	fmt.Printf("s[0] is %T = %d (%c)\n", s[0], s[0], s[0])

	// range decodes UTF-8: i is the byte index, ch is the rune.
	fmt.Println("\nrange over \"Go🚀\":")
	for i, ch := range "Go🚀" {
		fmt.Printf("  byte %d: %c (U+%04X)\n", i, ch, ch)
	}

	// Build strings with a Builder, not += in a loop (which is O(n²)).
	parts := []string{"Images", "Documents", "Videos"}
	var sb strings.Builder
	for i, p := range parts {
		if i > 0 {
			sb.WriteString(" · ")
		}
		sb.WriteString(p)
	}
	fmt.Printf("\nbuilt: %s\n", sb.String())
}
