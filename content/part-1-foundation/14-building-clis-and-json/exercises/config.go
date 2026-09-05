package exercises

// Config is a small settings type persisted as JSON, like gorg's rules file.
// Only exported fields are visible to encoding/json; the struct tags set the
// JSON keys, and omitempty drops DryRun from the output when it is false.
type Config struct {
	Root       string   `json:"root"`
	Categories []string `json:"categories"`
	DryRun     bool     `json:"dry_run,omitempty"`
}

// EncodeConfig returns c as human-readable (indented) JSON.
//
// TODO: use json.MarshalIndent with a two-space indent ("", "  ").
func EncodeConfig(c Config) ([]byte, error) {
	// TODO: implement
	return nil, nil
}

// DecodeConfig parses JSON bytes back into a Config.
//
// TODO: json.Unmarshal into a Config value and return it (with any error).
func DecodeConfig(data []byte) (Config, error) {
	// TODO: implement
	return Config{}, nil
}
