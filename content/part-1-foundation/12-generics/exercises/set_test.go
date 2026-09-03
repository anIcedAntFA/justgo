package exercises

import (
	"slices"
	"testing"
)

func TestSet(t *testing.T) {
	t.Skip("Chapter 12 exercise: implement Set, then delete this Skip")

	s := NewSet[string]()
	s.Add("a")
	s.Add("b")
	s.Add("a") // duplicate — no effect

	if got := s.Len(); got != 2 {
		t.Errorf("Len() = %d, want 2", got)
	}
	if !s.Contains("a") {
		t.Error(`Contains("a") = false, want true`)
	}
	if s.Contains("z") {
		t.Error(`Contains("z") = true, want false`)
	}

	s.Remove("a")
	if s.Contains("a") {
		t.Error(`after Remove("a"), Contains("a") = true, want false`)
	}
	if got := s.Len(); got != 1 {
		t.Errorf("after Remove, Len() = %d, want 1", got)
	}

	items := s.Items()
	slices.Sort(items) // Items() order is unspecified — sort before comparing
	if !slices.Equal(items, []string{"b"}) {
		t.Errorf("Items() = %v, want [b]", items)
	}
}

func TestSetInt(t *testing.T) {
	t.Skip("Chapter 12 exercise: implement Set, then delete this Skip")

	s := NewSet[int]()
	for _, n := range []int{3, 1, 4, 1, 5, 9, 2, 6, 5} {
		s.Add(n)
	}
	if got := s.Len(); got != 7 { // {3,1,4,5,9,2,6}
		t.Errorf("Len() = %d, want 7", got)
	}
}
