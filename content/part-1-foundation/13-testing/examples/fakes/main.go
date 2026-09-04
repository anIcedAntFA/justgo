// Command fakes shows why "accept interfaces" (Chapter 7) is what makes code
// testable. Greeter depends on the UserStore interface, so in main we wire a
// real (in-memory) store, and in main_test.go we inject a fake — no database,
// fast and deterministic.
//
// Run the demo:
//
//	go run .
//
// Run the tests (the point of the example):
//
//	go test -v
package main

import (
	"errors"
	"fmt"
)

// ErrUserNotFound is returned when a lookup misses.
var ErrUserNotFound = errors.New("user not found")

// User is the domain type.
type User struct {
	ID   int
	Name string
}

// UserStore is the dependency Greeter needs. Depending on this interface — not
// a concrete database — is the seam a test swaps out.
type UserStore interface {
	GetUser(id int) (*User, error)
}

// Greeter is the code under test.
type Greeter struct {
	store UserStore
}

// Greet builds a greeting for the user with the given id.
func (g *Greeter) Greet(id int) (string, error) {
	u, err := g.store.GetUser(id)
	if err != nil {
		return "", err
	}
	return "Hello, " + u.Name, nil
}

// memStore is a trivial real implementation for the demo's main.
type memStore struct{ users map[int]*User }

func (m *memStore) GetUser(id int) (*User, error) {
	u, ok := m.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return u, nil
}

func main() {
	g := &Greeter{store: &memStore{users: map[int]*User{1: {ID: 1, Name: "Alice"}}}}

	msg, err := g.Greet(1)
	fmt.Printf("Greet(1) = %q, err = %v\n", msg, err)

	_, err = g.Greet(99)
	fmt.Printf("Greet(99) err = %v\n", err)
}
