// Command subcommands is the Chapter 14 "subcommand dispatch" example: git-style
// subcommands (like `gorg stats` and `gorg undo`) built with only the flag
// package. Switch on the first argument, then give each subcommand its own
// FlagSet so their flags stay isolated.
//
// Run it from this directory:
//
//	go run . stats --json ./some/dir
//	go run . undo
//	go run .            (prints usage)
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	// os.Args[1] is the subcommand name; each handler parses os.Args[2:] itself.
	switch os.Args[1] {
	case "stats":
		stats(os.Args[2:])
	case "undo":
		undo(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand %q\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func stats(args []string) {
	fs := flag.NewFlagSet("stats", flag.ExitOnError)
	asJSON := fs.Bool("json", false, "emit stats as JSON")
	_ = fs.Parse(args)

	dir := "."
	if fs.NArg() > 0 {
		dir = fs.Arg(0)
	}
	fmt.Printf("stats for %s (json=%t)\n", dir, *asJSON)
}

func undo(args []string) {
	fs := flag.NewFlagSet("undo", flag.ExitOnError)
	_ = fs.Parse(args)
	fmt.Println("reverting the most recent run")
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: subcommands <stats|undo> [flags]")
}
