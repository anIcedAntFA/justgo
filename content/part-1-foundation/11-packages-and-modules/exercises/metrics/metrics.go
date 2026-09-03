// Package metrics is a leaf library for Exercise 2: it imports nothing from this
// project, so nothing here can ever start an import cycle. Package report (one layer
// up) imports metrics — a one-way arrow, the dependency direction Go forces you to
// keep acyclic.
package metrics

// Mean returns the average of xs, or 0 for an empty slice (a useful zero value —
// no panic, no error to handle).
//
// TODO: sum xs and divide by len(xs); return 0 when xs is empty.
func Mean(xs []float64) float64 {
	return 0 // TODO
}
