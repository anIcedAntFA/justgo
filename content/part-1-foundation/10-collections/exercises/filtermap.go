package exercises

// Exercise 2: Filter and Map with loops.
//
// Go has no built-in .filter()/.map() — you write the loop. Later (Chapter 12)
// generics let you write these once for any element type; here, do the int
// versions by hand so the mechanics are clear.

// Filter returns a new slice holding only the elements for which keep returns
// true, in their original order.
//
// TODO: range over s, append the ones keep says to keep.
func Filter(s []int, keep func(int) bool) []int {
	// TODO: implement
	return nil
}

// Map returns a new slice where each element is transform applied to the
// corresponding element of s (same length as s).
//
// TODO: preallocate make([]int, len(s)) and fill it by index.
func Map(s []int, transform func(int) int) []int {
	// TODO: implement
	return nil
}
