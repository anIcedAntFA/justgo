// Package exercises holds the coding exercises for Chapter 13: Testing.
//
// How to use: read the TODO, implement the function, then remove the t.Skip in
// the matching _test.go and run `go test ./...` until it passes. Reading the
// provided table-driven tests is half the lesson — that's the pattern you'll
// write yourself all through the course.
package exercises

// IsPalindrome reports whether s reads the same forwards and backwards,
// ignoring case, spaces, and punctuation — so "A man, a plan, a canal: Panama"
// is a palindrome. It must be correct for Unicode (compare by rune, not byte).
//
// TODO:
//  1. Walk s by rune (range over the string, or []rune(s)).
//  2. Keep only letters and digits (see unicode.IsLetter / unicode.IsDigit),
//     lower-casing each with unicode.ToLower.
//  3. Compare the filtered runes from both ends moving inward.
//
// The empty string and a single character are palindromes.
func IsPalindrome(s string) bool {
	return false // TODO
}
