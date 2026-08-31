package exercises

import (
	"errors"
	"testing"
)

func TestKVStoreSentinels(t *testing.T) {
	// t.Skip("Chapter 09 exercise: implement KVStore, then delete this Skip")

	s := NewKVStore()

	if err := s.Create("a", "apple"); err != nil {
		t.Fatalf("Create(a) unexpected error: %v", err)
	}

	// Create on an existing key → ErrKeyExists.
	if err := s.Create("a", "again"); !errors.Is(err, ErrKeyExists) {
		t.Errorf("Create(existing) = %v, want ErrKeyExists", err)
	}

	// Get hit → value, no error.
	got, err := s.Get("a")
	if err != nil || got != "apple" {
		t.Errorf("Get(a) = %q, %v; want \"apple\", nil", got, err)
	}

	// Get miss → ErrKeyNotFound.
	if _, err := s.Get("missing"); !errors.Is(err, ErrKeyNotFound) {
		t.Errorf("Get(missing) error = %v, want ErrKeyNotFound", err)
	}

	// Delete hit → nil, then a second delete misses.
	if err := s.Delete("a"); err != nil {
		t.Errorf("Delete(a) = %v, want nil", err)
	}
	if err := s.Delete("a"); !errors.Is(err, ErrKeyNotFound) {
		t.Errorf("Delete(a) second time = %v, want ErrKeyNotFound", err)
	}
}
