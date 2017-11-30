package main

import (
	cli "github.com/tumult-project/go-command"
	"github.com/tumult-project/tumult/command"
)

var (
	// Version string filled by the compiler.
	Version string
)

func main() {
	// TODO:
	// * ADD -version flag at the same level as -help
	//
	// versionPtr := flag.Bool("version", false, "Show version information")
	/*
		if *versionPtr {
			program := os.Args[0]
			v := version.GetVersion(Version)
			vnumber := v.VersionNumber()
			fmt.Printf("%s %s\n", program, vnumber)
			os.Exit(0)
		}
	*/

	cli.Register(command.VersionCommand)

	cli.Parse()
}
