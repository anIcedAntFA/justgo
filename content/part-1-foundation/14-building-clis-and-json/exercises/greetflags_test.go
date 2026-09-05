package exercises

import "testing"

func TestGreet(t *testing.T) {
	t.Skip("Chapter 14 exercise: implement Greet, then delete this Skip")

	cases := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{"default name", nil, "Hello, World!", false},
		{"named", []string{"-name", "Go"}, "Hello, Go!", false},
		{"loud", []string{"-name", "Go", "-loud"}, "HELLO, GO!", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Greet(tc.args)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Greet(%v) error = %v, wantErr %t", tc.args, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Greet(%v) = %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}
