package main

import (
	"flag"
	"fmt"
	"os"
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

func version() {
	version := "v0.0.0"
	fmt.Printf("%s %s\n", os.Args[0], version)
}

func main() {
	versionPtr := flag.Bool("version", false, "Show version information")

	flag.Usage = usage
	flag.Parse()

	if *versionPtr {
		version()
		os.Exit(0)
	}

	usage()
}
