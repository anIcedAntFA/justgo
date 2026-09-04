package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// assertFileContains is a helper: t.Helper() makes a failure point at the
// CALLER's line, not this one, so you can tell which assertion in the test
// actually broke.
func assertFileContains(t *testing.T, path, want string) {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	if !strings.Contains(string(data), want) {
		t.Errorf("%s does not contain %q; got:\n%s", path, want, data)
	}
}

func TestWriteConfig(t *testing.T) {
	// t.TempDir() gives a per-test directory that is removed automatically —
	// no cleanup to register, and parallel tests can't collide.
	path := filepath.Join(t.TempDir(), "app.conf")

	// t.Cleanup runs when the test finishes, pass or fail. Here it's just to
	// show the mechanism; TempDir already handles the directory itself.
	t.Cleanup(func() { _ = os.Remove(path) })

	if err := WriteConfig(path, map[string]string{"name": "gitm", "env": "test"}); err != nil {
		t.Fatalf("WriteConfig: %v", err)
	}

	assertFileContains(t, path, "name=gitm")
	assertFileContains(t, path, "env=test")
}
