package exercises

import "errors"

// Exercise 1: Sentinel Errors.
//
// Build a tiny in-memory key-value store whose failures are exported sentinel
// errors, so callers can branch on them with errors.Is instead of matching message
// strings.

// Exported sentinels — the ErrXxx naming convention.
var (
	ErrKeyNotFound = errors.New("key not found")
	ErrKeyExists   = errors.New("key already exists")
)

// KVStore is an in-memory string-to-string store.
type KVStore struct {
	data map[string]string
}

// NewKVStore returns a ready-to-use store.
//
// TODO: initialise the data map so Create can write into it.
func NewKVStore() *KVStore {
	// TODO: implement
	return nil
}

// Get returns the value for key, or ErrKeyNotFound if the key is absent.
//
// TODO: implement using the comma-ok map lookup.
func (s *KVStore) Get(key string) (string, error) {
	// TODO: implement
	return "", nil
}

// Create stores value under key. If the key already exists it returns ErrKeyExists
// and leaves the existing value untouched.
//
// TODO: implement.
func (s *KVStore) Create(key, value string) error {
	// TODO: implement
	return nil
}

// Delete removes key, or returns ErrKeyNotFound if it isn't present.
//
// TODO: implement.
func (s *KVStore) Delete(key string) error {
	// TODO: implement
	return nil
}
