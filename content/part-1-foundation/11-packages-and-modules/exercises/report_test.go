package exercises

import (
	"testing"

	"github.com/anIcedAntFA/justgo/content/part-1-foundation/11-packages-and-modules/exercises/report"
)

func TestReportSummary(t *testing.T) {
	t.Skip("Chapter 11 exercise: implement metrics.Mean and report.Summary, then delete this Skip")

	cases := []struct {
		name    string
		samples []float64
		want    string
	}{
		{"three values", []float64{2, 4, 6}, "samples=3 avg=4.00"},
		{"empty is zero", nil, "samples=0 avg=0.00"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := report.Summary(tc.samples); got != tc.want {
				t.Errorf("Summary(%v) = %q, want %q", tc.samples, got, tc.want)
			}
		})
	}
}
