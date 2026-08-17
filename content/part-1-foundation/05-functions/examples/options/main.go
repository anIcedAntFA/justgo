// Command options demonstrates the functional options pattern — Go's idiomatic
// answer to "configurable value with defaults" in the absence of default
// parameters. Each WithX returns a closure that mutates the config; the
// constructor starts from defaults and applies whatever options it's given.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"time"
)

type ServerConfig struct {
	Port    int
	Timeout time.Duration
	MaxConn int
}

// Option mutates a ServerConfig. It's a first-class function value; each WithX
// below is a closure over its argument.
type Option func(*ServerConfig)

func WithPort(port int) Option {
	return func(c *ServerConfig) { c.Port = port }
}

func WithTimeout(t time.Duration) Option {
	return func(c *ServerConfig) { c.Timeout = t }
}

func WithMaxConn(n int) Option {
	return func(c *ServerConfig) { c.MaxConn = n }
}

// NewServer builds a config from sensible defaults, then applies each override.
func NewServer(opts ...Option) *ServerConfig {
	config := &ServerConfig{
		Port:    8080,
		Timeout: 30 * time.Second,
		MaxConn: 100,
	}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

func main() {
	defaults := NewServer()
	fmt.Printf("defaults:  %+v\n", *defaults)

	custom := NewServer(
		WithPort(9000),
		WithTimeout(60*time.Second),
	)
	fmt.Printf("custom:    %+v\n", *custom) // MaxConn stays at the default 100
}
