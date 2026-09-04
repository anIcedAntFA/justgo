package exercises

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Skip("Chapter 13 exercise: implement IsPalindrome, then delete this Skip")

	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"empty", "", true},
		{"single char", "x", true},
		{"simple", "racecar", true},
		{"mixed case + punctuation", "A man, a plan, a canal: Panama", true},
		{"not a palindrome", "hello", false},
		{"unicode palindrome", "上海海上", true},
		{"unicode non-palindrome", "世界", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.in); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.in, got, tt.want)
			}
		})
	}
}
