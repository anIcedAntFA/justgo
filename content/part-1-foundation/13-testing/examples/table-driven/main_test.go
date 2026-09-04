package main

import (
	"errors"
	"fmt"
	"testing"
)

// TestClassify is the canonical shape: cases as data, one t.Run subtest each.
// Add a row to add a case — that's the whole appeal.
func TestClassify(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want string
	}{
		{"zero", 0, "zero"},
		{"negative", -4, "negative"},
		{"positive even", 6, "positive-even"},
		{"positive odd", 7, "positive-odd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Classify(tt.in); got != tt.want {
				t.Errorf("Classify(%d) = %q; want %q", tt.in, got, tt.want)
			}
		})
	}
}

// ExampleClassify is a third kind of test go test runs. It prints to stdout and
// the // Output: comment is verified — if Classify changes, this fails. It also
// shows up as documentation in `go doc` and on pkg.go.dev.
func ExampleClassify() {
	fmt.Println(Classify(0))
	fmt.Println(Classify(-4))
	fmt.Println(Classify(7))
	// Output:
	// zero
	// negative
	// positive-odd
}

// TestDivide folds an error expectation into the table. Note the two styles:
// wantErr is the specific sentinel checked with errors.Is (nil means "no
// error"), and we only check the value on the success path.
func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr error
	}{
		{"normal", 10, 2, 5, nil},
		{"negative result", -10, 2, -5, nil},
		{"divide by zero", 1, 0, 0, ErrDivideByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Divide(%v, %v) err = %v; want %v", tt.a, tt.b, err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return // error case — the value is meaningless, stop here
			}
			if got != tt.want {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
