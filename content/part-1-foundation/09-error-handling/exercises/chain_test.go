package exercises

import (
	"errors"
	"strings"
	"testing"
)

func TestWrappingChain(t *testing.T) {
	// t.Skip("Chapter 09 exercise: implement a, b, c, then delete this Skip")

	err := a()
	if err == nil {
		t.Fatal("a() = nil, want a wrapped error")
	}

	// The message shows every layer's context, in order.
	msg := err.Error()
	for _, part := range []string{"a:", "b:", "c:", "root cause"} {
		if !strings.Contains(msg, part) {
			t.Errorf("error %q missing segment %q", msg, part)
		}
	}

	// %w kept the sentinel detectable from the top of the chain.
	if !errors.Is(err, ErrRoot) {
		t.Errorf("errors.Is(a(), ErrRoot) = false, want true")
	}
}
