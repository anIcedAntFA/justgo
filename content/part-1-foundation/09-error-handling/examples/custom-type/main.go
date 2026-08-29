// Command custom-type shows a custom error type that carries structured data, an
// Unwrap method so the chain stays inspectable, and errors.As to extract the type
// and read its fields.
package main

import (
	"errors"
	"fmt"
)

// HTTPError carries structured data (a status code and URL) alongside the wrapped
// underlying error — more than a sentinel's single message can express.
type HTTPError struct {
	StatusCode int
	URL        string
	Err        error
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d from %s: %v", e.StatusCode, e.URL, e.Err)
}

// Unwrap lets errors.Is/As traverse into the wrapped Err.
func (e *HTTPError) Unwrap() error { return e.Err }

// errRateLimited is the underlying sentinel a caller might still want to detect.
var errRateLimited = errors.New("rate limited")

func fetch(url string) error {
	// Simulate a 429 response wrapping a sentinel.
	return fmt.Errorf("fetch failed: %w", &HTTPError{
		StatusCode: 429,
		URL:        url,
		Err:        errRateLimited,
	})
}

func main() {
	err := fetch("https://api.example.com/data")

	// errors.As extracts the *HTTPError from the chain so we can read its fields.
	// The pre-1.26 form: declare a target, pass its address, check the bool.
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		fmt.Printf("via As:     status=%d url=%s\n", httpErr.StatusCode, httpErr.URL)
	}

	// errors.AsType (Go 1.26) is the type-safe generic form — the type is a
	// parameter and the value is returned, comma-ok style. No any target to misuse.
	if he, ok := errors.AsType[*HTTPError](err); ok {
		fmt.Printf("via AsType: status=%d url=%s\n", he.StatusCode, he.URL)
	}
	// via As:     status=429 url=https://api.example.com/data
	// via AsType: status=429 url=https://api.example.com/data

	// Unwrap means the inner sentinel is still detectable through both wrappings.
	fmt.Println("rate limited:", errors.Is(err, errRateLimited)) // true
}
