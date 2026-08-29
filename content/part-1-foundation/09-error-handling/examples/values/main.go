// Command values shows the foundation of Go error handling: an error is an
// ordinary value returned last, checked with `if err != nil`. It also contrasts
// errors.New (static message) with fmt.Errorf (dynamic message).
package main

import (
	"errors"
	"fmt"
)

// errNegative is a static condition — errors.New is the right tool.
var errNegative = errors.New("age cannot be negative")

// validate returns nil on success, a non-nil error on failure. The error is the
// last return value — the Go convention.
func validate(age int) error {
	if age < 0 {
		return errNegative // static message: errors.New value
	}
	if age > 150 {
		// dynamic message: fmt.Errorf interpolates the offending value
		return fmt.Errorf("implausible age: %d (must be <= 150)", age)
	}
	return nil
}

func main() {
	for _, age := range []int{30, -1, 200} {
		if err := validate(age); err != nil {
			fmt.Printf("age %d rejected: %v\n", age, err)
			continue
		}
		fmt.Printf("age %d ok\n", age)
	}
	// age 30 ok
	// age -1 rejected: age cannot be negative
	// age 200 rejected: implausible age: 200 (must be <= 150)
}
