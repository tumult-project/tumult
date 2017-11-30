package command

import (
	"fmt"
	"os"

	"github.com/tumult-project/go-command"
	"github.com/tumult-project/tumult/version"
)

// VersionCommand is a command prints version information
var VersionCommand = &command.Command{
	Name:  "version",
	Usage: "Show version information",
	Run:   run,
}

// Run executes the command
func run(cmd *command.Command, args []string) {
	program := args[0]
	v := version.GetVersion("9.9.9")
	vnumber := v.VersionNumber()
	fmt.Printf("%s %s\n", program, vnumber)
	os.Exit(0)
}
