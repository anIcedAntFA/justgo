// Command satisfaction demonstrates implicit interface satisfaction and
// polymorphism: Rectangle and Circle satisfy Shape just by having the right
// methods — no "implements" keyword — so a []Shape can hold both.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"math"
)

// Shape is satisfied by anything with Area and Perimeter methods.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func main() {
	// Rectangle and Circle both satisfy Shape automatically.
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 3},
	}

	for _, s := range shapes {
		fmt.Printf("%-12T area=%7.2f perimeter=%7.2f\n", s, s.Area(), s.Perimeter())
	}
}
