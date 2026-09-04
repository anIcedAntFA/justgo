package main

import "testing"

const benchInput = "The quick brown fox jumps over the lazy dog — héllo, 世界!"

// A correctness test still matters: the benchmark below would happily report
// ReverseBytes as "faster" even though it mangles Unicode. Speed is not
// correctness.
func TestReverseRunesRoundTrip(t *testing.T) {
	if got := ReverseRunes(ReverseRunes(benchInput)); got != benchInput {
		t.Errorf("ReverseRunes twice = %q; want original", got)
	}
}

// BenchmarkReverseBytes and BenchmarkReverseRunes use the Go 1.24+ b.Loop()
// form. It's cleaner than the old for i := 0; i < b.N; i++ and stops the
// compiler from optimising the call away. Compare with:
//
//	go test -bench=. -benchmem
func BenchmarkReverseBytes(b *testing.B) {
	for b.Loop() {
		ReverseBytes(benchInput)
	}
}

func BenchmarkReverseRunes(b *testing.B) {
	for b.Loop() {
		ReverseRunes(benchInput)
	}
}
