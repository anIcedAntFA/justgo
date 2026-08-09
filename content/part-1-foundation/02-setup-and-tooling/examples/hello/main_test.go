package main

import "testing"

// TestGoAge is a preview of Go's built-in testing — no framework, no config.
// Proper testing is covered in Chapter 13.
func TestGoAge(t *testing.T) {
	want := 19
	if got := 2026 - 2007; got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
