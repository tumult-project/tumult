package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tumult-project/tumult/version"
)

var (
	// Version string filled by the compiler.
	Version string
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-version] [-help] <command> [<args>]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Available commands are:\n")
	fmt.Fprintf(os.Stderr, "  version")
	fmt.Fprintf(os.Stderr, "\tShow version information\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Available flags are:\n")
	fmt.Fprintf(os.Stderr, "  -help\n")
	fmt.Fprintf(os.Stderr, "\tShow this help message\n")
	flag.PrintDefaults()
}

func main() {
	versionPtr := flag.Bool("version", false, "Show version information")

	versionCommand := flag.NewFlagSet("version", flag.ExitOnError)

	flag.Usage = usage
	flag.Parse()

	if *versionPtr {
		program := os.Args[0]
		v := version.GetVersion(Version)
		vnumber := v.VersionNumber()
		fmt.Printf("%s %s\n", program, vnumber)
		os.Exit(0)
	}

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	//
	// TODO:
	// * Check for unkown commands
	switch os.Args[1] {
	case "version":
		versionCommand.Parse(os.Args[2:])
	default:
		usage()
		os.Exit(1)
	}

	if versionCommand.Parsed() {
		fmt.Println("version command")
	}
}
