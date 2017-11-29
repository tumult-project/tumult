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
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Available flags are:\n")
	fmt.Fprintf(os.Stderr, "  -help\n")
	fmt.Fprintf(os.Stderr, "\tShow this help message\n")
	flag.PrintDefaults()
}

func main() {
	versionPtr := flag.Bool("version", false, "Show version information")

	flag.Usage = usage
	flag.Parse()

	if *versionPtr {
		program := os.Args[0]
		v := version.GetVersion(Version)
		vnumber := v.VersionNumber()
		fmt.Printf("%s %s\n", program, vnumber)
		os.Exit(0)
	}

	usage()
}
