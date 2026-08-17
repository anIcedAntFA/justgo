package exercises

import "testing"

func TestLimiter(t *testing.T) {
	t.Skip("Chapter 05 exercise: implement Limiter, then delete this Skip")

	cases := []struct {
		name  string
		max   int
		calls int // total calls to make
		// wantErr[i] is whether call i (0-indexed) should return an error
	}{
		{"allows three", 3, 5},
		{"allows one", 1, 3},
		{"allows none", 0, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			allow := Limiter(tc.max)
			for i := range tc.calls {
				err := allow()
				wantErr := i >= tc.max
				if wantErr && err == nil {
					t.Errorf("call %d: got nil, want an error (max=%d)", i+1, tc.max)
				}
				if !wantErr && err != nil {
					t.Errorf("call %d: got error %v, want nil (max=%d)", i+1, err, tc.max)
				}
			}
		})
	}
}
