package exercises

import (
	"errors"
	"testing"
)

func TestSafeExecute(t *testing.T) {
	// t.Skip("Chapter 09 exercise: implement safeExecute, then delete this Skip")

	errBoom := errors.New("boom")

	cases := []struct {
		name    string
		fn      func() error
		wantErr string // "" means expect nil
	}{
		{"success", func() error { return nil }, ""},
		{"returns error", func() error { return errBoom }, "boom"},
		{"panics", func() error { panic("something exploded") }, "recovered from panic: something exploded"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := safeExecute(tc.fn)
			if tc.wantErr == "" {
				if err != nil {
					t.Errorf("safeExecute = %v, want nil", err)
				}
				return
			}
			if err == nil || err.Error() != tc.wantErr {
				t.Errorf("safeExecute = %v, want %q", err, tc.wantErr)
			}
		})
	}
}
