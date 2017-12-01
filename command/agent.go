package command

import (
	"log"

	command "github.com/tumult-project/go-command"
	"github.com/tumult-project/tumult/agent"
)

// AgentCommand is a command starts an agent
var AgentCommand = &command.Command{
	Name:  "agent",
	Usage: "Runs an agent",
	Run: func(cmd *command.Command, args []string) {
		agent, err := agent.NewAgent()
		if err != nil {
			log.Fatal("can't start agent", err)
		}
		agent.Start()
	},
}
