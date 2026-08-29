// Command join shows errors.Join (Go 1.20+) bundling several errors into one value
// whose tree errors.Is can still walk against each joined error — plus a custom
// type doing the same by implementing the multi-error Unwrap() []error.
package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyName   = errors.New("name is empty")
	ErrNegativeAge = errors.New("age is negative")
)

type user struct {
	Name string
	Age  int
}

// validate accumulates every problem, then joins them. errors.Join skips nil and
// returns nil if the slice is empty — so a valid user yields no error.
func validate(u user) error {
	var errs []error
	if u.Name == "" {
		errs = append(errs, ErrEmptyName)
	}
	if u.Age < 0 {
		errs = append(errs, ErrNegativeAge)
	}
	return errors.Join(errs...)
}

// ValidationErrors is a custom multi-error: it carries the failures AND stays
// inspectable, because Unwrap() []error lets errors.Is walk into each one.
type ValidationErrors struct {
	Errs []error
}

func (e *ValidationErrors) Error() string {
	msgs := make([]string, len(e.Errs))
	for i, err := range e.Errs {
		msgs[i] = err.Error()
	}
	return "validation: " + strings.Join(msgs, "; ")
}

// Unwrap returns the slice form (Go 1.20+), so errors.Is/As/AsType see every error.
func (e *ValidationErrors) Unwrap() []error { return e.Errs }

func main() {
	err := validate(user{Name: "", Age: -1})
	fmt.Println(err) // each joined error prints on its own line:
	// name is empty
	// age is negative

	// Both joined errors remain individually detectable.
	fmt.Println("empty name:", errors.Is(err, ErrEmptyName))     // true
	fmt.Println("negative age:", errors.Is(err, ErrNegativeAge)) // true

	fmt.Println("valid user err:", validate(user{Name: "Ada", Age: 36})) // <nil>

	// A custom multi-error is just as inspectable, via Unwrap() []error.
	custom := &ValidationErrors{Errs: []error{ErrEmptyName, ErrNegativeAge}}
	fmt.Println("custom:", custom)                                         // validation: name is empty; age is negative
	fmt.Println("custom has empty name:", errors.Is(custom, ErrEmptyName)) // true
}
