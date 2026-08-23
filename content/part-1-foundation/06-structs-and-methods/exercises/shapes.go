package exercises

// Shape is the embedded base. Circle and Square each embed it, so Shape's
// Describe method is PROMOTED onto them — you can call circle.Describe() even
// though Describe is defined here.
type Shape struct {
	Name string
}

// Describe returns a string of the form "I am a <Name>", e.g. "I am a circle".
//
// TODO: implement with fmt.Sprintf and add the import yourself.
func (s Shape) Describe() string {
	return "" // TODO: replace
}

// Circle embeds Shape and adds a radius.
type Circle struct {
	Shape
	Radius float64
}

// Area returns the circle's area (π·r²).
//
// TODO: implement using math.Pi and add the import yourself.
func (c Circle) Area() float64 {
	return 0 // TODO: replace
}

// Square embeds Shape and adds a side length.
type Square struct {
	Shape
	Side float64
}

// Area returns the square's area (side²).
//
// TODO: implement.
func (s Square) Area() float64 {
	return 0 // TODO: replace
}
