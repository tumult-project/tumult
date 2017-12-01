package command

import (
	"fmt"

	"github.com/tumult-project/go-command"
	"github.com/tumult-project/tumult/version"
)

// VersionCommand is a command prints version information
var VersionCommand = &command.Command{
	Name:  "version",
	Usage: "Show version information",
	Run: func(cmd *command.Command, args []string) {
		program := args[0]
		v := version.GetVersion()
		vnumber := v.VersionNumber()
		fmt.Printf("%s %s\n", program, vnumber)
	},
}
