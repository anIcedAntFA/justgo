// Command helpers-and-cleanup demonstrates test plumbing that isn't about
// assertions: a t.Helper() assertion helper, t.TempDir() for a throwaway
// directory, and t.Cleanup() for teardown. The function under test writes a
// tiny key=value config file.
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
	"fmt"
	"os"
	"path/filepath"
)

// WriteConfig writes "key=value\n" lines to path, creating the file. It returns
// an error rather than panicking so tests can exercise the failure path.
func WriteConfig(path string, kv map[string]string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	// Named return + deferred close reports a flush/close error the write loop
	// wouldn't otherwise see — the idiomatic way to not silently drop it.
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for k, v := range kv {
		if _, err := fmt.Fprintf(f, "%s=%s\n", k, v); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	dir, err := os.MkdirTemp("", "cfg-*")
	if err != nil {
		panic(err)
	}
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "app.conf")
	if err := WriteConfig(path, map[string]string{"name": "gitm"}); err != nil {
		panic(err)
	}

	data, _ := os.ReadFile(path)
	fmt.Printf("wrote %s:\n%s", path, data)
}
