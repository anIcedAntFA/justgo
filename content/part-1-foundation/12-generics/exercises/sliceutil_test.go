package exercises

import (
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	t.Skip("Chapter 12 exercise: implement Map, then delete this Skip")

	nums := []int{1, 2, 3, 4, 5}
	got := Map(nums, func(n int) int { return n * 2 })
	want := []int{2, 4, 6, 8, 10}
	if !slices.Equal(got, want) {
		t.Errorf("Map doubled = %v, want %v", got, want)
	}

	// T and U differ here: []int -> []string.
	labels := Map(nums, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	wantLabels := []string{"odd", "even", "odd", "even", "odd"}
	if !slices.Equal(labels, wantLabels) {
		t.Errorf("Map labels = %v, want %v", labels, wantLabels)
	}
}

func TestFilter(t *testing.T) {
	t.Skip("Chapter 12 exercise: implement Filter, then delete this Skip")

	nums := []int{1, 2, 3, 4, 5}
	got := Filter(nums, func(n int) bool { return n%2 == 0 })
	want := []int{2, 4}
	if !slices.Equal(got, want) {
		t.Errorf("Filter evens = %v, want %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	t.Skip("Chapter 12 exercise: implement Reduce, then delete this Skip")

	nums := []int{1, 2, 3, 4, 5}
	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	if sum != 15 {
		t.Errorf("Reduce sum = %d, want 15", sum)
	}

	// U differs from T: fold []int into a string.
	joined := Reduce(nums, "", func(acc string, n int) string {
		if acc == "" {
			return string(rune('0' + n))
		}
		return acc + "-" + string(rune('0'+n))
	})
	if joined != "1-2-3-4-5" {
		t.Errorf("Reduce joined = %q, want %q", joined, "1-2-3-4-5")
	}
}
