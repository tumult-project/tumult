package command

import (
	"flag"
	"fmt"
	"os"
)

// Command is a command implementation
type Command struct {
	Name  string
	Usage string
	Run   func(cmd *Command, args []string)
}

var (
	commands []*Command
)

// Usage prints usage information to stderr
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-version] [-help] <command> [<args>]\n", os.Args[0])
	if len(commands) > 0 {
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Available commands are:\n")
		for _, cmd := range commands {
			fmt.Fprintf(os.Stderr, "  %s\t%s\n", cmd.Name, cmd.Usage)
		}
	}
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Available flags are:\n")
	fmt.Fprintf(os.Stderr, "  -help\n")
	fmt.Fprintf(os.Stderr, "\tShow this help message\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
}

// Register registers a Command
func Register(c *Command) {
	commands = append(commands, c)
}

// Parse parses and executes commands given its options and arguments
func Parse() {
	flag.Usage = Usage
	flag.Parse()

	// Verify that a command has been provided
	// flag.Args() will be the list of arguments
	// flag.Arg(0) will be the command
	if len(flag.Args()) < 1 {
		Usage()
		os.Exit(2)
	}

	// TODO:
	// * Check for unkown commands
	for _, cmd := range commands {
		if cmd.Name == flag.Arg(0) {
			cmd.Run(cmd, flag.Args())
		}
	}
}
