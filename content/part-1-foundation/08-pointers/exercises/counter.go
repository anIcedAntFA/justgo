package exercises

// Counter drills receiver semantics. Increment needs a POINTER receiver so its
// change persists on the original; Value only reads, so a value receiver is
// enough. (In production you'd normally make every method on a type with any
// pointer receiver a pointer receiver too, for consistency — here the split is
// deliberate, to show the difference.)
type Counter struct {
	count int
}

// Increment adds one to the counter. It must be a pointer receiver: a value
// receiver would increment a copy and the change would not persist.
//
// TODO: implement.
func (c *Counter) Increment() {
	// TODO: implement
}

// Value returns the current count.
//
// TODO: implement.
func (c Counter) Value() int {
	// TODO: implement
	return 0
}
