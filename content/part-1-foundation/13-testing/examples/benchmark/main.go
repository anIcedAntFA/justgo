// Command benchmark compares two string-reversal implementations and shows the
// modern benchmark loop (for b.Loop(), Go 1.24+). One reverses bytes (wrong for
// multi-byte runes), one reverses runes (correct for Unicode) — a reminder that
// a benchmark measures speed, never correctness.
//
// Run the demo:
//
//	go run .
//
// Run the benchmarks (the point of the example):
//
//	go test -bench=. -benchmem
package main

import "fmt"

// ReverseBytes reverses the raw bytes. Fast, but it corrupts any multi-byte
// UTF-8 character (e.g. "世界").
func ReverseBytes(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// ReverseRunes reverses by rune, so it's correct for Unicode. It allocates a
// []rune, so it's a touch slower and heavier — the trade the benchmark exposes.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	const s = "héllo, 世界"
	fmt.Printf("input:        %q\n", s)
	fmt.Printf("ReverseBytes: %q  (corrupts multi-byte runes)\n", ReverseBytes(s))
	fmt.Printf("ReverseRunes: %q  (correct)\n", ReverseRunes(s))
}
