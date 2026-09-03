package exercises

import (
	"errors"
	"testing"
)

func TestResultOk(t *testing.T) {
	t.Skip("Chapter 12 exercise: implement Result, then delete this Skip")

	r := Ok(42)
	if !r.IsOk() {
		t.Error("IsOk() = false, want true")
	}
	if got := r.Unwrap(); got != 42 {
		t.Errorf("Unwrap() = %d, want 42", got)
	}
	if got := r.UnwrapOr(0); got != 42 {
		t.Errorf("UnwrapOr(0) = %d, want 42", got)
	}
}

func TestResultErr(t *testing.T) {
	t.Skip("Chapter 12 exercise: implement Result, then delete this Skip")

	r := Err[int](errors.New("boom"))
	if r.IsOk() {
		t.Error("IsOk() = true, want false")
	}
	if got := r.UnwrapOr(-1); got != -1 {
		t.Errorf("UnwrapOr(-1) = %d, want -1 (fallback on error)", got)
	}

	// Unwrap on an error Result must panic.
	defer func() {
		if recover() == nil {
			t.Error("Unwrap() on an error Result did not panic")
		}
	}()
	_ = r.Unwrap()
}
