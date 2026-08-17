package exercises

// Compose returns a new function that applies g first, then f — i.e. it computes
// f(g(x)). Both f and g are first-class function values.
//
//	doubleThenAddOne := Compose(addOne, double)
//	doubleThenAddOne(5) // (5*2)+1 = 11
//
// TODO: return a func(int) int that calls g then f.
func Compose(f, g func(int) int) func(int) int {
	return nil // TODO: replace
}
