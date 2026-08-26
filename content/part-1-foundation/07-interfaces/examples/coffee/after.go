package main

// AFTER — one small interface.
//
// A Brewer is anything that can turn Grounds into a Cup. That is the single seam
// that actually varies. Each brew method becomes its own type that satisfies
// Brewer implicitly — no "implements" keyword.
type Brewer interface {
	Brew(g Grounds) Cup
}

// Phin is the Vietnamese drip filter: slow, strong, no moving parts.
type Phin struct{}

func (Phin) Brew(g Grounds) Cup {
	return Cup{Drink: "Cà phê phin", Origin: g.Origin, Strength: 8}
}

// EspressoMachine forces hot water through the grounds under pressure.
type EspressoMachine struct{}

func (EspressoMachine) Brew(g Grounds) Cup {
	return Cup{Drink: "Espresso", Origin: g.Origin, Strength: 9}
}

// ColdBrew steeps grounds in cold water for many hours; more hours, smoother and
// a touch stronger. It carries state (Hours), which a string method never could.
type ColdBrew struct {
	Hours int
}

func (c ColdBrew) Brew(g Grounds) Cup {
	strength := 4
	if c.Hours >= 12 {
		strength = 6
	}
	return Cup{Drink: "Cold brew", Origin: g.Origin, Strength: strength}
}

// Serve is the whole payoff. It accepts a Brewer (any of them) and never has to
// change when a new brew method appears — contrast brewBAD, which you must
// reopen and edit. "Accept interfaces, return structs": it takes the Brewer
// interface and returns a concrete Cup.
func Serve(b Brewer, g Grounds) Cup {
	return b.Brew(g)
}
