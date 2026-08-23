// Command builder shows a NewXxx constructor that sets defaults, plus method
// chaining (the builder pattern). Chaining only works because the setters use
// pointer receivers and return the same *HTTPClient.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"time"
)

type HTTPClient struct {
	baseURL string
	timeout time.Duration
	retries int
}

// NewHTTPClient is the constructor convention: validate/seed defaults, return a
// pointer.
func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		baseURL: baseURL,
		timeout: 30 * time.Second, // default
		retries: 3,                // default
	}
}

// SetTimeout returns the receiver so calls can be chained.
func (c *HTTPClient) SetTimeout(d time.Duration) *HTTPClient {
	c.timeout = d
	return c
}

func (c *HTTPClient) SetRetries(n int) *HTTPClient {
	c.retries = n
	return c
}

func main() {
	defaults := NewHTTPClient("https://api.example.com")
	fmt.Printf("defaults: %+v\n", *defaults)

	custom := NewHTTPClient("https://api.example.com").
		SetTimeout(60 * time.Second).
		SetRetries(5)
	fmt.Printf("custom:   %+v\n", *custom)
}
