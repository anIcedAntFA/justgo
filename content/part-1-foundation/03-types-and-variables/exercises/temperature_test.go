package exercises

import "testing"

// Table-driven tests — the idiomatic Go pattern you'll see all through this
// course (Chapter 13 covers it in full). Cases use exact float64 values so
// there is no floating-point rounding to worry about here.

func TestCtoF(t *testing.T) {
	cases := []struct {
		name string
		in   Celsius
		want Fahrenheit
	}{
		{"freezing point", 0, 32},
		{"boiling point", 100, 212},
		{"equal point", -40, -40},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CtoF(tc.in); got != tc.want {
				t.Errorf("CtoF(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestFtoC(t *testing.T) {
	cases := []struct {
		name string
		in   Fahrenheit
		want Celsius
	}{
		{"freezing point", 32, 0},
		{"boiling point", 212, 100},
		{"equal point", -40, -40},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FtoC(tc.in); got != tc.want {
				t.Errorf("FtoC(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
