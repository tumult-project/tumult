package main

import (
	"github.com/tumult-project/tumult/command"
)

var (
	// Version string filled by the compiler.
	Version string
)

func main() {
	//versionPtr := flag.Bool("version", false, "Show version information")
	/*
		if *versionPtr {
			program := os.Args[0]
			v := version.GetVersion(Version)
			vnumber := v.VersionNumber()
			fmt.Printf("%s %s\n", program, vnumber)
			os.Exit(0)
		}
	*/

	command.Register(command.VersionCommand)

	command.Parse()
}
