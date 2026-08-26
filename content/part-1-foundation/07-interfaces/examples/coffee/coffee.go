// Command coffee tells one story twice: brewing coffee WITHOUT an interface
// (before.go — a switch on a string) and WITH one (after.go — a Brewer
// interface). main.go plays out a café scene so you can feel the difference.
//
// Roasting and grinding are the same for every cup, so they stay concrete
// functions. Only *how you brew* varies from cup to cup — that is the one seam
// that earns an interface. (See the README: "extract interfaces only when you
// actually need polymorphism".)
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

// Origin is where the bean comes from. It is a plain typed constant, not a
// separate type per origin — origin is data, not behaviour, so it does not need
// its own type (and modelling it as one would only tempt a needless type switch).
type Origin string

const (
	Robusta Origin = "Robusta"
	Arabica Origin = "Arabica"
)

// RoastLevel and GrindSize are likewise just labelled values.
type RoastLevel string

const (
	LightRoast  RoastLevel = "light"
	MediumRoast RoastLevel = "medium"
	DarkRoast   RoastLevel = "dark"
)

type GrindSize string

const (
	Coarse GrindSize = "coarse"
	Fine   GrindSize = "fine"
)

// Bean is a raw green coffee bean, before roasting.
type Bean struct {
	Origin Origin
}

// RoastedBean is what Roast produces.
type RoastedBean struct {
	Origin Origin
	Level  RoastLevel
}

// Grounds is roasted coffee, ground and ready to brew.
type Grounds struct {
	Origin Origin
	Level  RoastLevel
	Size   GrindSize
}

// Cup is the finished drink handed to the customer.
type Cup struct {
	Drink    string // e.g. "Cà phê phin", "Espresso", "Cold brew"
	Origin   Origin
	Strength int // 1..10, just to make different brew methods visibly differ
}

func (c Cup) String() string {
	return fmt.Sprintf("%s (%s, strength %d/10)", c.Drink, c.Origin, c.Strength)
}

// Roast turns a green Bean into a RoastedBean. Concrete — there is exactly one
// way to roast here, so no interface.
func Roast(b Bean, level RoastLevel) RoastedBean {
	return RoastedBean{Origin: b.Origin, Level: level}
}

// Grind turns a RoastedBean into Grounds. Concrete for the same reason.
func Grind(rb RoastedBean, size GrindSize) Grounds {
	return Grounds{Origin: rb.Origin, Level: rb.Level, Size: size}
}
