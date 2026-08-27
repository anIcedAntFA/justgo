// Package exercises holds the coding exercises for Chapter 07: Interfaces.
//
// How to use: read the TODO, implement the code, then remove the t.Skip in the
// matching _test.go and run `go test ./...` until it passes.
package exercises

import "math"

// Shape is satisfied by any type with Area and Perimeter methods. Rectangle,
// Circle, and Triangle must each satisfy it — implicitly, just by having the
// methods. No "implements" keyword.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle has a width and a height.
type Rectangle struct {
	Width, Height float64
}

// TODO: implement Area (Width*Height) and Perimeter (2*(Width+Height)).
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

// Circle has a radius.
type Circle struct {
	Radius float64
}

// TODO: implement Area (π·r²) and Perimeter (2·π·r) using math.Pi — add the
// math import yourself.
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle is defined by the lengths of its three sides.
type Triangle struct {
	A, B, C float64
}

// TODO: implement Perimeter (A+B+C) and Area via Heron's formula:
//
//	s = (A+B+C)/2
//	area = √(s(s-A)(s-B)(s-C))
//
// math.Sqrt lives in the math package.
func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	area := math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
	return area
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// TotalArea sums the Area of every shape in the slice — the payoff of the
// interface: one function over a mixed slice of concrete types.
//
// TODO: implement.
func TotalArea(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}
