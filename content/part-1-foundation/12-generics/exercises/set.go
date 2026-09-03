package exercises

// Set is a generic set of unique values. T is constrained to comparable because
// the elements are used as map keys. The backing map uses struct{} values,
// which occupy zero memory — the classic "set from a map" idiom.
//
// TODO: give Set a field, e.g. `m map[T]struct{}`.
type Set[T comparable] struct {
	// TODO: add the backing map field
}

// NewSet builds an empty, ready-to-use Set[T].
//
// TODO: return a *Set[T] with its map initialized (make(map[T]struct{})).
func NewSet[T comparable]() *Set[T] {
	return nil // TODO
}

// Add inserts v. Adding a value already present is a no-op.
//
// TODO: set s.m[v] = struct{}{}.
func (s *Set[T]) Add(v T) {
	// TODO
}

// Remove deletes v. Removing an absent value is a no-op.
//
// TODO: delete(s.m, v).
func (s *Set[T]) Remove(v T) {
	// TODO
}

// Contains reports whether v is in the set.
//
// TODO: use the comma-ok map read.
func (s *Set[T]) Contains(v T) bool {
	return false // TODO
}

// Len returns the number of elements.
//
// TODO: return len(s.m).
func (s *Set[T]) Len() int {
	return 0 // TODO
}

// Items returns the set's elements as a slice, in unspecified order.
//
// TODO: collect the map keys into a []T.
func (s *Set[T]) Items() []T {
	return nil // TODO
}
