// Command receivers shows the single most important struct concept: a value
// receiver operates on a COPY (the original is untouched), while a pointer
// receiver operates on the ORIGINAL (mutations stick).
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// ScaleCopy has a VALUE receiver — it mutates a copy, so the caller's Rectangle
// is left unchanged. staticcheck even flags these writes as "ineffective"
// (SA4005) — the linter catching the exact footgun this demo is about, which is
// why we deliberately silence it here.
//
//nolint:staticcheck // intentional: value receiver mutates a copy — that's the lesson
func (r Rectangle) ScaleCopy(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// ScaleReal has a POINTER receiver — it mutates the original through the pointer.
func (r *Rectangle) ScaleReal(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	rect.ScaleCopy(2)
	fmt.Printf("after ScaleCopy(2):  %+v  // unchanged — value receiver got a copy\n", rect)

	// rect is addressable (a variable), so Go rewrites this as (&rect).ScaleReal(2).
	rect.ScaleReal(2)
	fmt.Printf("after ScaleReal(2):  %+v  // changed — pointer receiver hit the original\n", rect)
}
