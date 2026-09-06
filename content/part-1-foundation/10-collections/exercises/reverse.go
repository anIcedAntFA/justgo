// Package exercises holds the coding exercises for Chapter 10: Collections.
//
// How to use: read the TODO, implement the function, then remove the t.Skip in
// the matching _test.go and run `go test ./...` until it passes.
package exercises

// Exercise 1: Reverse a slice into a new one.
//
// Return a reversed COPY of s. The original must be left untouched — this is the
// discipline that keeps slice aliasing bugs away. Do it without the slices
// package (write the loop yourself).
//
// TODO: allocate a result with make([]int, len(s)) and fill it back-to-front.
func Reverse(s []int) []int {
	// rev := slices.Clone(s)
	rev := make([]int, len(s))

	// if len(rev) == 0 {
	// 	return rev
	// }
	// if len(rev) == 1 {
	// 	return rev
	// }

	for i, v := range s {
		rev[len(rev)-i-1] = v
	}

	return rev
}

// original = []int{1, 2, 3, 4}
//
// [] => []
// [4] => [4]
// [4, 3, 2, 1]
//
// input:
// index:  0   1   2   3
// 				 ↓   ↓   ↓   ↓
// s:     [1] [2] [3] [4]
//
// output:
// index:   0   1   2   3
//          ↓   ↓   ↓   ↓
// result: [4] [3] [2] [1]
//
// mapping:
// s[0] → result[3]
// s[1] → result[2]
// s[2] → result[1]
// s[3] → result[0]
//
// if input index i
// => output index = len(s) - i - 1
//
// test:
// i = 0 -> 4 - 0 - 1 = 3
// i = 1 -> 4 - 1 - 1 = 2
// i = 2 -> 4 - 2 - 1 = 1
// i = 3 -> 4 - 3 - 1 = 0
