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
	return &HTTPClient{
		baseURL: baseURL,
		timeout: 30 * time.Second,
		retries: 3,
	}
}

// SetTimeout overrides the timeout and returns the receiver so calls chain.
//
// TODO: set c.timeout and return c.
func (c *HTTPClient) SetTimeout(d time.Duration) *HTTPClient {
	c.timeout = d
	return c
}

// SetRetries overrides the retry count and returns the receiver so calls chain.
//
// TODO: set c.retries and return c.
func (c *HTTPClient) SetRetries(n int) *HTTPClient {
	c.retries = n
	return c
}
