// Package exercises holds the coding exercises for Chapter 07: Interfaces.
//
// How to use: read the TODO, implement the code, then remove the t.Skip in the
// matching _test.go and run `go test ./...` until it passes.
package exercises

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
func (r Rectangle) Area() float64      { return 0 }
func (r Rectangle) Perimeter() float64 { return 0 }

// Circle has a radius.
type Circle struct {
	Radius float64
}

// TODO: implement Area (π·r²) and Perimeter (2·π·r) using math.Pi — add the
// math import yourself.
func (c Circle) Area() float64      { return 0 }
func (c Circle) Perimeter() float64 { return 0 }

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
func (t Triangle) Area() float64      { return 0 }
func (t Triangle) Perimeter() float64 { return 0 }

// TotalArea sums the Area of every shape in the slice — the payoff of the
// interface: one function over a mixed slice of concrete types.
//
// TODO: implement.
func TotalArea(shapes []Shape) float64 {
	return 0 // TODO: replace
}
