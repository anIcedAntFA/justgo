// Package bank is a tiny demo library showing Go's visibility rule: capitalization,
// not an `export` keyword, decides what other packages can see. The main package
// next door imports bank and can reach only the capitalized identifiers here.
package bank

// Account is exported (capital A) — other packages can name this type.
type Account struct {
	Owner   string // exported field — readable from outside package bank
	balance int    // unexported field — private to package bank
}

// Open is exported: the constructor other packages call. It is named Open, not
// NewAccount, to avoid stutter — callers write bank.Open, which already reads well.
func Open(owner string, initial int) *Account {
	return &Account{Owner: owner, balance: initial}
}

// Balance is an exported method — the only way outsiders can read the private
// balance field.
func (a *Account) Balance() int { return a.balance }

// Deposit is exported; it mutates the unexported field through validated logic, so
// the invariant (no negative deposits) lives inside the package that owns the data.
func (a *Account) Deposit(amount int) {
	if !positive(amount) {
		return
	}
	a.balance += amount
}

// positive is unexported — an internal helper other packages cannot call.
func positive(n int) bool { return n > 0 }
