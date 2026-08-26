package exercises

// This exercise is the payoff of "accept interfaces": because Serve takes a
// Brewer interface (not a concrete espresso machine), you can test it by passing
// a fake Brewer — no real hardware, no I/O. That is exactly how interfaces make
// code testable, the theme you will lean on hard from Chapter 11 onward.

// Grounds is roasted, ground coffee ready to brew. Kept minimal on purpose.
type Grounds struct {
	Origin string
}

// Cup is the finished drink.
type Cup struct {
	Drink    string
	Strength int
}

// Brewer is anything that can turn Grounds into a Cup.
type Brewer interface {
	Brew(g Grounds) Cup
}

// Serve runs a brewer over the grounds and returns the cup. It is deliberately
// tiny — the point of the exercise is not Serve's body but that you can test it
// with a fake Brewer.
//
// TODO: implement Serve (one line: call b.Brew(g) and return it).
func Serve(b Brewer, g Grounds) Cup {
	return Cup{} // TODO: replace
}
