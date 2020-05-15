// cmd_completion.go: Helpers for completion

package subcommands

import (
	"fmt"
)

// CommandsList is a structure which implements the built-in help subcommand
type CommandList struct {

	// We embed the NoFlags option, because we accept no
	// command-line flags.
	NoFlags
}

// Execute the commands-command.
func (c *CommandList) Execute(args []string) int {
	//
	// Examine each known sub-command
	//
	for _, c := range known {

		//
		// Get the name of the subcommand.
		//
		name, _ := c.Info()

		fmt.Printf("%s\n", name)
	}
	return 0
}

// Info returns the name of this subcommand, along with a one-line synopsis.
func (h *CommandList) Info() (string, string) {
	return "commands", "Show all commands, useful for bash-completion."
}
