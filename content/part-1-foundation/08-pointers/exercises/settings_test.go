package exercises

import "testing"

// intPtr and boolPtr build addressable pointers for the table below — the point
// of the exercise is that a non-nil pointer to a zero value differs from nil.
func intPtr(v int) *int    { return &v }
func boolPtr(v bool) *bool { return &v }

func TestResolve(t *testing.T) {
	t.Skip("Chapter 08 exercise: implement Resolve, then delete this Skip")

	cases := []struct {
		name string
		in   Settings
		want Resolved
	}{
		{"all nil uses defaults", Settings{}, Resolved{Timeout: 30, Verbose: false}},
		{"timeout overridden", Settings{Timeout: intPtr(5)}, Resolved{Timeout: 5, Verbose: false}},
		{"timeout set to zero, not default", Settings{Timeout: intPtr(0)}, Resolved{Timeout: 0, Verbose: false}},
		{"verbose set to true", Settings{Verbose: boolPtr(true)}, Resolved{Timeout: 30, Verbose: true}},
		{"both set", Settings{Timeout: intPtr(90), Verbose: boolPtr(true)}, Resolved{Timeout: 90, Verbose: true}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Resolve(tc.in); got != tc.want {
				t.Errorf("Resolve(%+v) = %+v, want %+v", tc.in, got, tc.want)
			}
		})
	}
}
