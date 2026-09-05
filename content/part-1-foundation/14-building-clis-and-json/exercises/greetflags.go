// Package exercises holds the coding exercises for Chapter 14: Building CLIs & JSON.
//
// How to use: read the TODO, implement the function, then remove the t.Skip in
// the matching _test.go and run `go test ./...` until it passes.
package exercises

// Greet parses CLI-style arguments and returns a greeting.
//
// TODO: build a flag.FlagSet (use flag.NewFlagSet("greet", flag.ContinueOnError))
// with a -name flag (default "World") and a -loud bool flag. Parse args, then
// return "Hello, <name>!" — upper-cased (strings.ToUpper) when -loud is set.
// Return any parse error to the caller instead of exiting.
//
// Examples:
//
//	Greet(nil)                          → "Hello, World!"
//	Greet([]string{"-name", "Go"})      → "Hello, Go!"
//	Greet([]string{"-name","Go","-loud"}) → "HELLO, GO!"
func Greet(args []string) (string, error) {
	// TODO: implement
	return "", nil
}
