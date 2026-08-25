package exercises

import (
	"testing"
	"time"
)

func TestNewHTTPClientDefaults(t *testing.T) {
	// t.Skip("Chapter 06 exercise: implement NewHTTPClient, then delete this Skip")

	c := NewHTTPClient("https://api.example.com")
	if c == nil {
		t.Fatal("NewHTTPClient returned nil")
	}
	if c.baseURL != "https://api.example.com" {
		t.Errorf("baseURL = %q, want %q", c.baseURL, "https://api.example.com")
	}
	if c.timeout != 30*time.Second {
		t.Errorf("timeout = %v, want %v", c.timeout, 30*time.Second)
	}
	if c.retries != 3 {
		t.Errorf("retries = %d, want 3", c.retries)
	}
}

func TestHTTPClientChaining(t *testing.T) {
	// t.Skip("Chapter 06 exercise: implement the setters, then delete this Skip")

	c := NewHTTPClient("https://api.example.com").
		SetTimeout(60 * time.Second).
		SetRetries(5)

	if c == nil {
		t.Fatal("chained calls returned nil")
	}
	if c.timeout != 60*time.Second {
		t.Errorf("timeout = %v, want %v", c.timeout, 60*time.Second)
	}
	if c.retries != 5 {
		t.Errorf("retries = %d, want 5", c.retries)
	}
}
