// Command wrapping shows %w building a chain of context, and how %w preserves the
// original error for errors.Is while %v would discard it.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// loadConfig wraps the OS error with %w, adding context but keeping the original
// inspectable.
func loadConfig(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("loadConfig(%s): %w", path, err)
	}
	return nil
}

// initApp wraps again — each layer adds its own context to the chain.
func initApp(path string) error {
	if err := loadConfig(path); err != nil {
		return fmt.Errorf("initApp: %w", err)
	}
	return nil
}

func main() {
	err := initApp("/no/such/config.yaml")

	// The full chain reads outermost-first, original error preserved at the end.
	fmt.Println(err)
	// initApp: loadConfig(/no/such/config.yaml): open /no/such/config.yaml: no such file or directory

	// %w kept fs.ErrNotExist inspectable through two layers of wrapping.
	fmt.Println("is ErrNotExist:", errors.Is(err, fs.ErrNotExist)) // true

	// Compare: %v flattens to text and loses the wrapped error.
	flattened := fmt.Errorf("initApp: %v", err)
	fmt.Println("still detectable after string flattening:", errors.Is(flattened, fs.ErrNotExist)) // false
}
