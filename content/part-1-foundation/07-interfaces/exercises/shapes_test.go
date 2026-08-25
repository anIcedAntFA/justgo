package exercises

import (
	"math"
	"testing"
)

func TestShapeAreaPerimeter(t *testing.T) {
	t.Skip("Chapter 07 exercise: implement the Shape types, then delete this Skip")

	cases := []struct {
		name      string
		shape     Shape
		wantArea  float64
		wantPerim float64
	}{
		{"rectangle", Rectangle{Width: 10, Height: 5}, 50, 30},
		{"circle r=2", Circle{Radius: 2}, 4 * math.Pi, 4 * math.Pi},
		{"triangle 3-4-5", Triangle{A: 3, B: 4, C: 5}, 6, 12},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.shape.Area(); math.Abs(got-tc.wantArea) > 1e-9 {
				t.Errorf("Area() = %v, want %v", got, tc.wantArea)
			}
			if got := tc.shape.Perimeter(); math.Abs(got-tc.wantPerim) > 1e-9 {
				t.Errorf("Perimeter() = %v, want %v", got, tc.wantPerim)
			}
		})
	}
}

func TestTotalArea(t *testing.T) {
	t.Skip("Chapter 07 exercise: implement TotalArea, then delete this Skip")

	shapes := []Shape{
		Rectangle{Width: 10, Height: 5}, // 50
		Circle{Radius: 2},               // 4π
		Triangle{A: 3, B: 4, C: 5},      // 6
	}
	want := 50 + 4*math.Pi + 6
	if got := TotalArea(shapes); math.Abs(got-want) > 1e-9 {
		t.Errorf("TotalArea() = %v, want %v", got, want)
	}
}
