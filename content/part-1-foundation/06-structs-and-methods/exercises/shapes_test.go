package exercises

import (
	"math"
	"testing"
)

func TestShapeDescribePromotion(t *testing.T) {
	t.Skip("Chapter 06 exercise: implement Shape.Describe, then delete this Skip")

	// Describe is defined on Shape but called through the embedding types — this
	// is method promotion.
	c := Circle{Shape: Shape{Name: "circle"}, Radius: 2}
	s := Square{Shape: Shape{Name: "square"}, Side: 3}

	if got, want := c.Describe(), "I am a circle"; got != want {
		t.Errorf("Circle.Describe() = %q, want %q", got, want)
	}
	if got, want := s.Describe(), "I am a square"; got != want {
		t.Errorf("Square.Describe() = %q, want %q", got, want)
	}
}

func TestShapeArea(t *testing.T) {
	t.Skip("Chapter 06 exercise: implement Area, then delete this Skip")

	cases := []struct {
		name string
		got  float64
		want float64
	}{
		{"circle r=2", Circle{Radius: 2}.Area(), 4 * math.Pi},
		{"circle r=1", Circle{Radius: 1}.Area(), math.Pi},
		{"square s=3", Square{Side: 3}.Area(), 9},
		{"square s=5", Square{Side: 5}.Area(), 25},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if math.Abs(tc.got-tc.want) > 1e-9 {
				t.Errorf("Area() = %v, want %v", tc.got, tc.want)
			}
		})
	}
}
