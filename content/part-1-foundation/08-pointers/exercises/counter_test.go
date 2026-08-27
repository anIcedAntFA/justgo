package exercises

import "testing"

func TestCounterIncrement(t *testing.T) {
	t.Skip("Chapter 08 exercise: implement Counter.Increment and Value, then delete this Skip")

	cases := []struct {
		name string
		n    int // how many times to Increment
		want int
	}{
		{"no increments", 0, 0},
		{"one", 1, 1},
		{"several", 5, 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var c Counter
			for range tc.n {
				c.Increment() // addressable var — Go takes &c automatically
			}
			if got := c.Value(); got != tc.want {
				t.Errorf("after %d Increment() calls, Value() = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
