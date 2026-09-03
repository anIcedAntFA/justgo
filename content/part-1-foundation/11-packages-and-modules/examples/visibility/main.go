// Command visibility shows how capitalization controls cross-package access. This
// main package imports the sibling bank package and can touch only its exported
// (capitalized) identifiers — the unexported ones are compile errors, shown below.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"

	"github.com/anIcedAntFA/justgo/content/part-1-foundation/11-packages-and-modules/examples/visibility/bank"
)

func main() {
	acc := bank.Open("Alice", 100) // ✅ Open is exported
	acc.Deposit(50)                // ✅ Deposit is exported
	acc.Deposit(-999)              // ignored: bank's own unexported guard rejects it

	fmt.Println("owner:  ", acc.Owner)     // ✅ Owner field is exported
	fmt.Println("balance:", acc.Balance()) // ✅ Balance method is exported

	// Every line below is a COMPILE error — it reaches for unexported names that
	// live behind package bank's boundary. Uncomment one to watch the compiler
	// enforce visibility:
	//
	//	_ = acc.balance      // ❌ balance is unexported (lowercase b)
	//	_ = bank.positive(1) // ❌ positive is unexported
}
