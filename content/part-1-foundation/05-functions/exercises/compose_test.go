package exercises

import "testing"

func TestCompose(t *testing.T) {
	t.Skip("Chapter 05 exercise: implement Compose, then delete this Skip")

	addOne := func(n int) int { return n + 1 }
	double := func(n int) int { return n * 2 }
	square := func(n int) int { return n * n }

	cases := []struct {
		name string
		f, g func(int) int
		in   int
		want int
	}{
		{"double then addOne", addOne, double, 5, 11}, // (5*2)+1
		{"addOne then double", double, addOne, 5, 12}, // (5+1)*2
		{"square then addOne", addOne, square, 3, 10}, // (3*3)+1
		{"order matters", square, addOne, 3, 16},      // (3+1)^2
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Compose(tc.f, tc.g)(tc.in)
			if got != tc.want {
				t.Errorf("Compose(...)(%d) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
