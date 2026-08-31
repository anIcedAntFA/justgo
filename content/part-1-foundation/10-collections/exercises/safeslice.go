package exercises

// Exercise 5: A subslice you can safely append to.
//
// Plain s[low:high] shares the backing array with s and keeps spare capacity, so
// appending to it can silently overwrite s. Return a subslice of s[low:high]
// that is SAFE to append to — appending to the result must never modify s.
//
// Two idiomatic fixes (either passes the test):
//   - slices.Clone(s[low:high]) — copy into a fresh backing array.
//   - the three-index full slice expression s[low:high:high] — caps the result so
//     the next append is forced to reallocate.
//
// TODO: return an independent view of s[low:high].
func SafeSubslice(s []int, low, high int) []int {
	// TODO: implement
	return nil
}
