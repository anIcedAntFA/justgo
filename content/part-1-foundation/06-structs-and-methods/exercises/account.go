// Package exercises holds the coding exercises for Chapter 06: Structs & Methods.
//
// How to use: read the TODO, implement the code, then remove the t.Skip in the
// matching _test.go and run `go test ./...` until it passes.
package exercises

// BankAccount models a simple account. Deposit and Withdraw must MUTATE the
// balance, which is the whole reason they take a pointer receiver — a value
// receiver would change only a copy and the balance would never move.
type BankAccount struct {
	Owner   string
	Balance float64
}

// Deposit adds amount to the balance. It must return a non-nil error (and leave
// the balance untouched) when amount is negative.
//
// TODO: implement. Use errors.New or fmt.Errorf and add the import yourself.
func (a *BankAccount) Deposit(amount float64) error {
	return nil // TODO: replace
}

// Withdraw subtracts amount from the balance. It must return a non-nil error
// (and leave the balance untouched) when amount is negative or exceeds the
// current balance.
//
// TODO: implement.
func (a *BankAccount) Withdraw(amount float64) error {
	return nil // TODO: replace
}

// String returns a human-readable summary, e.g. "Alice: 150.00". Implementing
// String makes BankAccount satisfy fmt.Stringer, so fmt prints it with this.
//
// TODO: implement with fmt.Sprintf.
func (a *BankAccount) String() string {
	return "" // TODO: replace
}
