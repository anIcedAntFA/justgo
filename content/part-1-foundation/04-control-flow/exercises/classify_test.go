package exercises

import "testing"

func TestIsPrime(t *testing.T) {
	t.Skip("Chapter 04 exercise: implement IsPrime, then delete this Skip")

	cases := []struct {
		in   int
		want bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{9, false},
		{17, true},
		{25, false},
		{97, true},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			if got := IsPrime(tc.in); got != tc.want {
				t.Errorf("IsPrime(%d) = %t, want %t", tc.in, got, tc.want)
			}
		})
	}
}

func TestClassify(t *testing.T) {
	t.Skip("Chapter 04 exercise: implement Classify, then delete this Skip")

	cases := []struct {
		name string
		in   int
		want string
	}{
		{"fizzbuzz fifteen", 15, "fizzbuzz"},
		{"fizzbuzz forty-five", 45, "fizzbuzz"},
		{"prime beats even", 2, "prime"},
		{"prime three", 3, "prime"},
		{"prime seven", 7, "prime"},
		{"even four", 4, "even"},
		{"even eight", 8, "even"},
		{"even six", 6, "even"},
		{"odd nine", 9, "odd"},
		{"odd one", 1, "odd"},
		{"odd twenty-five", 25, "odd"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Classify(tc.in); got != tc.want {
				t.Errorf("Classify(%d) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
