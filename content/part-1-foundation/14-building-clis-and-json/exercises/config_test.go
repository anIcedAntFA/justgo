package exercises

import "testing"

func TestConfigRoundTrip(t *testing.T) {
	t.Skip("Chapter 14 exercise: implement EncodeConfig/DecodeConfig, then delete this Skip")

	in := Config{
		Root:       "/home/me/Downloads",
		Categories: []string{"Images", "Documents"},
		DryRun:     true,
	}

	data, err := EncodeConfig(in)
	if err != nil {
		t.Fatalf("EncodeConfig: %v", err)
	}

	got, err := DecodeConfig(data)
	if err != nil {
		t.Fatalf("DecodeConfig: %v", err)
	}

	if got.Root != in.Root || got.DryRun != in.DryRun || len(got.Categories) != len(in.Categories) {
		t.Errorf("round-trip mismatch:\n got  %+v\n want %+v", got, in)
	}
}
