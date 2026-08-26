package exercises

import "testing"

// fakeBrewer is a test double: it satisfies Brewer without brewing anything
// real. It records whether it was called and with which grounds, and returns a
// canned Cup. This is the whole reason Serve accepts an interface — you can swap
// a real machine for a fake in a test.
type fakeBrewer struct {
	called bool
	gotIn  Grounds
	out    Cup
}

func (f *fakeBrewer) Brew(g Grounds) Cup {
	f.called = true
	f.gotIn = g
	return f.out
}

func TestServeUsesTheBrewer(t *testing.T) {
	t.Skip("Chapter 07 exercise: implement Serve, then delete this Skip")

	cases := []struct {
		name    string
		grounds Grounds
		want    Cup
	}{
		{"phin-like", Grounds{Origin: "Robusta"}, Cup{Drink: "Cà phê phin", Strength: 8}},
		{"cold-brew-like", Grounds{Origin: "Arabica"}, Cup{Drink: "Cold brew", Strength: 5}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fake := &fakeBrewer{out: tc.want}

			got := Serve(fake, tc.grounds)

			if !fake.called {
				t.Fatal("Serve did not call the brewer's Brew method")
			}
			if fake.gotIn != tc.grounds {
				t.Errorf("brewer received grounds %+v, want %+v", fake.gotIn, tc.grounds)
			}
			if got != tc.want {
				t.Errorf("Serve returned %+v, want %+v", got, tc.want)
			}
		})
	}
}
