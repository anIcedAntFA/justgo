package exercises

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Skip("Chapter 04 exercise: implement FizzBuzz, then delete this Skip")

	cases := []struct {
		name string
		in   int
		want string
	}{
		{"one", 1, "1"},
		{"two", 2, "2"},
		{"fizz three", 3, "Fizz"},
		{"buzz five", 5, "Buzz"},
		{"fizz six", 6, "Fizz"},
		{"seven", 7, "7"},
		{"buzz ten", 10, "Buzz"},
		{"fizzbuzz fifteen", 15, "FizzBuzz"},
		{"fizzbuzz thirty", 30, "FizzBuzz"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FizzBuzz(tc.in); got != tc.want {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
