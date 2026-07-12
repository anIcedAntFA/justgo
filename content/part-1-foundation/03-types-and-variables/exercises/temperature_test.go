package exercises

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	t.Skip("Chapter 03 exercise: implement CelsiusToFahrenheit, then delete this Skip")

	// Table-driven test — the idiomatic Go pattern you'll see all through this
	// course (Chapter 13 covers it in full). Cases use exact float64 values so
	// there is no floating-point rounding to worry about here.
	cases := []struct {
		name string
		in   float64
		want float64
	}{
		{"freezing point", 0, 32},
		{"boiling point", 100, 212},
		{"equal point", -40, -40},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CelsiusToFahrenheit(tc.in)
			if got != tc.want {
				t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
