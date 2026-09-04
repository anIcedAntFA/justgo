package main

import (
	"errors"
	"testing"
)

// fakeStore is a hand-written fake: a working, minimal UserStore for tests. No
// database, no network — just a map. For simple cases this is clearer than a
// generated mock.
type fakeStore struct {
	users map[int]*User
}

func (f *fakeStore) GetUser(id int) (*User, error) {
	u, ok := f.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return u, nil
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    string
		wantErr error
	}{
		{"known user", 1, "Hello, Alice", nil},
		{"unknown user", 99, "", ErrUserNotFound},
	}

	// Inject the fake once; each subtest exercises a different lookup.
	g := &Greeter{store: &fakeStore{users: map[int]*User{1: {ID: 1, Name: "Alice"}}}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := g.Greet(tt.id)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Greet(%d) err = %v; want %v", tt.id, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Greet(%d) = %q; want %q", tt.id, got, tt.want)
			}
		})
	}
}
