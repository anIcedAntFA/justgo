package exercises

import (
	"fmt"
	"testing"
)

func TestDurationString(t *testing.T) {
	t.Skip("Chapter 07 exercise: implement Duration.String, then delete this Skip")

	cases := []struct {
		in   Duration
		want string
	}{
		{3661, "1h 1m 1s"},
		{90, "1m 30s"},
		{45, "45s"},
		{3600, "1h 0m 0s"},
		{0, "0s"},
	}

	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			if got := tc.in.String(); got != tc.want {
				t.Errorf("Duration(%d).String() = %q, want %q", int(tc.in), got, tc.want)
			}
			// fmt must pick up String() through the Stringer interface.
			if got := fmt.Sprintf("%v", tc.in); got != tc.want {
				t.Errorf("fmt %%v = %q, want %q", got, tc.want)
			}
		})
	}
}
