package exercises

import "testing"

func TestDescribe(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want string
	}{
		{"string", "hi", `"hi"`},
		{"int", 42, "42"},
		{"float64", 3.5, "3.5"},
		{"bool true", true, "yes"},
		{"bool false", false, "no"},
		{"slice", []any{1, "x"}, `[1, "x"]`},
		{"map single key", map[string]any{"k": 1}, "{k: 1}"},
		{"default uint", uint(5), "5 (uint)"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Describe(tc.in); got != tc.want {
				t.Errorf("Describe(%v) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
