package main

import "fmt"

// Order pairs a customer with the brew method they asked for. The Brewer field
// is an interface, so an order can hold ANY brew method — including ones this
// package has never heard of (see HomeMokaPot below).
type Order struct {
	Customer string
	Brewer   Brewer
}

func main() {
	// The shop roasts and grinds the same way for everyone.
	beans := Bean{Origin: Robusta}
	grounds := Grind(Roast(beans, DarkRoast), Fine)

	// --- BEFORE: the string-keyed switch ---
	fmt.Println("=== BEFORE (switch on a string) ===")
	fmt.Println(brewBAD("phin", grounds))
	fmt.Println(brewBAD("espresso", grounds))
	// A typo compiles fine and quietly produces nonsense at run time:
	fmt.Println(brewBAD("expreso", grounds)) // note the missing 's'

	// --- AFTER: the Brewer interface ---
	fmt.Println("\n=== AFTER (Brewer interface) ===")
	orders := []Order{
		{Customer: "Bà Tám", Brewer: Phin{}},
		{Customer: "John", Brewer: ColdBrew{Hours: 12}},
		{Customer: "Lan", Brewer: EspressoMachine{}},
	}
	for _, o := range orders {
		fmt.Printf("%-8s → %v\n", o.Customer, Serve(o.Brewer, grounds))
	}

	// The payoff: a customer brings gear this package was never written for.
	// HomeMokaPot satisfies Brewer just by having a Brew method — implicit
	// satisfaction. Serve does not change one line to accommodate it.
	fmt.Println("\n=== A customer brings their own gear ===")
	walkIn := Order{Customer: "Minh", Brewer: HomeMokaPot{}}
	fmt.Printf("%-8s → %v\n", walkIn.Customer, Serve(walkIn.Brewer, grounds))

	// Compile-time safety, for contrast:
	//   brewBAD("expreso", grounds) // compiles, wrong at run time
	//   Serve(Espres(), grounds)    // a typo like this is a COMPILE error
}

// HomeMokaPot stands in for a third party's type — imagine it living in someone
// else's package. It never imports this file, yet it satisfies Brewer simply by
// having Brew(Grounds) Cup. This is what "satisfy interfaces you don't own"
// buys you.
type HomeMokaPot struct{}

func (HomeMokaPot) Brew(g Grounds) Cup {
	return Cup{Drink: "Moka pot", Origin: g.Origin, Strength: 7}
}
