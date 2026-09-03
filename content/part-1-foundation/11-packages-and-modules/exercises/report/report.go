// Package report sits one layer above metrics and imports it. The dependency runs
// one way only: report -> metrics. metrics must never import report back, or the
// compiler rejects the cycle.
//
// Note: this stub ships with no imports. Part of the exercise is writing the import
// line for the metrics package yourself, using its full module path.
package report

// Summary formats a one-line report of the sample's average. It reuses metrics.Mean
// through the package's qualified name — you import your own package by its full
// module path, then call members as metrics.Mean.
//
// TODO: import "fmt" and the metrics package, then return
// fmt.Sprintf("samples=%d avg=%.2f", len(samples), metrics.Mean(samples)).
func Summary(samples []float64) string {
	return "" // TODO
}
