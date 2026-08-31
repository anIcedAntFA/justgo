package exercises

import (
	"slices"
	"testing"
)

func TestSafeSubslice(t *testing.T) {
	t.Skip("Chapter 10 exercise: implement SafeSubslice, then delete this Skip")

	// The heart of the exercise: appending to the returned slice must NOT reach
	// back into the parent's backing array.
	t.Run("append does not corrupt parent", func(t *testing.T) {
		parent := []int{1, 2, 3, 4, 5}
		before := slices.Clone(parent)

		sub := SafeSubslice(parent, 0, 2) // [1 2]
		sub = append(sub, 99)             // must not write into parent[2]

		if !slices.Equal(parent, before) {
			t.Errorf("parent corrupted: got %v, want %v", parent, before)
		}
		if want := []int{1, 2, 99}; !slices.Equal(sub, want) {
			t.Errorf("sub = %v, want %v", sub, want)
		}
	})

	t.Run("returns the right window", func(t *testing.T) {
		parent := []int{10, 20, 30, 40}
		if got, want := SafeSubslice(parent, 1, 3), []int{20, 30}; !slices.Equal(got, want) {
			t.Errorf("SafeSubslice(%v, 1, 3) = %v, want %v", parent, got, want)
		}
	})
}
