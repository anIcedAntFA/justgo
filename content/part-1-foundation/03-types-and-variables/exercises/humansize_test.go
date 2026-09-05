package exercises

import "testing"

func TestHumanSize(t *testing.T) {
	t.Skip("Chapter 03 exercise: implement HumanSize, then delete this Skip")

	cases := []struct {
		name string
		in   int64
		want string
	}{
		{"zero", 0, "0 B"},
		{"bytes", 512, "512 B"},
		{"exactly 1 KB", 1024, "1.0 KB"},
		{"one and a half KB", 1536, "1.5 KB"},
		{"one MB", 1048576, "1.0 MB"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := HumanSize(tc.in); got != tc.want {
				t.Errorf("HumanSize(%d) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
