package exercises

import "time"

// HTTPClient is configured through a constructor plus chainable setters. The
// fields are unexported — the constructor and setters are the only intended way
// to build and adjust one.
type HTTPClient struct {
	baseURL string
	timeout time.Duration
	retries int
}

// NewHTTPClient builds an HTTPClient with sensible defaults: timeout 30s,
// retries 3.
//
// TODO: return &HTTPClient{...} with baseURL set and the two defaults.
func NewHTTPClient(baseURL string) *HTTPClient {
	return nil // TODO: replace
}

// SetTimeout overrides the timeout and returns the receiver so calls chain.
//
// TODO: set c.timeout and return c.
func (c *HTTPClient) SetTimeout(d time.Duration) *HTTPClient {
	return nil // TODO: replace
}

// SetRetries overrides the retry count and returns the receiver so calls chain.
//
// TODO: set c.retries and return c.
func (c *HTTPClient) SetRetries(n int) *HTTPClient {
	return nil // TODO: replace
}
