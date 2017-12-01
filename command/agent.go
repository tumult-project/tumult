package command

import (
	"fmt"

	command "github.com/tumult-project/go-command"
)

// AgentCommand is a command starts tumult in server mode
var AgentCommand = &command.Command{
	Name:  "agent",
	Usage: "Runs an agent",
	Run: func(cmd *command.Command, args []string) {
		fmt.Println("Start agent!")
	},
}
