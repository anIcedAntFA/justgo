// Command sentinel shows exported ErrXxx sentinel values and how callers branch on
// them with errors.Is — even after the error has been wrapped with context.
package main

import (
	"errors"
	"fmt"
)

// Exported sentinels callers can check for. The ErrXxx naming is the convention.
var (
	ErrNotFound = errors.New("item not found")
	ErrConflict = errors.New("item already exists")
)

type store struct{ items map[string]string }

// get returns the sentinel ErrNotFound, wrapped with context. errors.Is still
// finds the sentinel through the wrapping.
func (s *store) get(id string) (string, error) {
	v, ok := s.items[id]
	if !ok {
		return "", fmt.Errorf("store.get %q: %w", id, ErrNotFound)
	}
	return v, nil
}

func main() {
	s := &store{items: map[string]string{"a": "apple"}}

	for _, id := range []string{"a", "z"} {
		v, err := s.get(id)
		switch {
		case errors.Is(err, ErrNotFound):
			fmt.Printf("%q: not found (would return 404)\n", id)
		case err != nil:
			fmt.Printf("%q: unexpected error: %v\n", id, err)
		default:
			fmt.Printf("%q: %s\n", id, v)
		}
	}
	// "a": apple
	// "z": not found (would return 404)
}
