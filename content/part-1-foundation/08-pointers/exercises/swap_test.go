package exercises

import "testing"

func TestSwap(t *testing.T) {
	t.Skip("Chapter 08 exercise: implement swap, then delete this Skip")

	cases := []struct {
		name         string
		a, b         int
		wantA, wantB int
	}{
		{"basic", 10, 20, 20, 10},
		{"equal", 5, 5, 5, 5},
		{"negatives", -1, 3, 3, -1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a, b := tc.a, tc.b
			swap(&a, &b)
			if a != tc.wantA || b != tc.wantB {
				t.Errorf("swap(&%d, &%d) => %d, %d; want %d, %d",
					tc.a, tc.b, a, b, tc.wantA, tc.wantB)
			}
		})
	}
}
