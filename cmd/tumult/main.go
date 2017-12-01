package main

import (
	cli "github.com/tumult-project/go-command"
	"github.com/tumult-project/tumult/command"
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

	var rootCmd = &cli.RootCommand{}

	rootCmd.AddCommand(command.VersionCommand)

	rootCmd.Parse()

	rootCmd.Usage()
}
