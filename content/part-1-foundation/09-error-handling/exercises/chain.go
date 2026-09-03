package exercises

import (
	"errors"
	"fmt"
)

// Exercise 2: Error Wrapping Chain.
//
// Three functions call each other: a → b → c. c fails with the sentinel ErrRoot.
// Each layer wraps the error with %w and adds its own context, so the final message
// shows the whole call path AND errors.Is can still detect ErrRoot from the top.

// ErrRoot is the original failure c produces.
var ErrRoot = errors.New("root cause")

// c is the innermost function — it fails with the sentinel.
//
// TODO: return ErrRoot wrapped with context "c: ..." using fmt.Errorf and %w.
func c() error {
	return fmt.Errorf("c: %w", ErrRoot)
}

// b calls c and wraps any error with its own context.
//
// TODO: if c() fails, return it wrapped as "b: %w".
func b() error {
	return fmt.Errorf("b: %w", c())
}

// a calls b and wraps any error with its own context.
//
// TODO: if b() fails, return it wrapped as "a: %w".
func a() error {
	return fmt.Errorf("a: %w", b())
}
