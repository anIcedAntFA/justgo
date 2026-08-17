package exercises

import "testing"

func TestDivide(t *testing.T) {
	t.Skip("Chapter 05 exercise: implement Divide, then delete this Skip")

	cases := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"simple", 10, 2, 5, false},
		{"fraction", 3, 4, 0.75, false},
		{"negative", -10, 5, -2, false},
		{"zero numerator", 0, 7, 0, false},
		{"divide by zero", 10, 0, 0, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.a, tc.b)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("Divide(%v, %v) = %v, nil; want an error", tc.a, tc.b, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("Divide(%v, %v) unexpected error: %v", tc.a, tc.b, err)
			}
			if got != tc.want {
				t.Errorf("Divide(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
