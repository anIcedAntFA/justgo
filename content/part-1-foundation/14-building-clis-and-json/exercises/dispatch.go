package exercises

import "errors"

// ErrUnknownCommand is returned by Dispatch for a name it does not recognize.
var ErrUnknownCommand = errors.New("unknown command")

// Dispatch maps a subcommand name to the action gorg would run. An empty name
// defaults to "organize" (the default command).
//
// TODO: return the action for each known name — "" and "organize" → "organize";
// "stats" → "stats"; "undo" → "undo"; "rules" → "rules". For anything else,
// return "" and an error that wraps ErrUnknownCommand (use fmt.Errorf with %w).
func Dispatch(name string) (string, error) {
	// TODO: implement
	return "", nil
}
