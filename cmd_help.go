// cmd_help.go: Implementation of the default help-command.

package subcommands

import (
	"flag"
	"fmt"
	"os"
	"path"
)

// Help is a structure which implements the built-in help subcommand
type Help struct {

	// We embed the NoFlags option, because we accept no
	// command-line flags.
	NoFlags
}

// Execute the help command.
//
// If there are no arguments we just show the available subcommands,
// if we're given a command then we show the flags that the subcommand
// accepts.
func (h *Help) Execute(args []string) {

	//
	// No flags?  Just dump commands.
	//
	if len(args) == 0 {

		fmt.Printf("Available subcommands:\n\n")
		dump()

		fmt.Printf("\nFor more details please run '%s help subcommand'.\n",
			path.Base(os.Args[0]))
		return
	}

	//
	// For each argument we were given.
	//
	for _, cmd := range args {

		//
		// Examine each known sub-command
		//
		for _, c := range known {

			//
			// Get the name of the subcommand.
			//
			name := c.Name()

			//
			// If the user specified this command.
			//
			if cmd == name {

				//
				// Create a new flagset using that name.
				//
				set := flag.NewFlagSet(name, flag.ExitOnError)
				c.Arguments(set)

				//
				// Count the flags this method accepts.
				//
				count := 0
				set.VisitAll(func(f *flag.Flag) {
					count++
				})

				//
				// If there are no flags show that.
				//
				if count == 0 {
					fmt.Printf("Synopsis:\n\t%s\n", c.Synopsis())
					fmt.Printf("\n")
					fmt.Printf("Usage:\n\t%s %s\n\n", path.Base(os.Args[0]), name)

				} else {

					//
					// Otherwise show each flag.
					//
					fmt.Printf("Synopsis:\n\t%s\n", c.Synopsis())
					fmt.Printf("\n")
					fmt.Printf("Usage:\n\t%s %s [flags]\n\n", path.Base(os.Args[0]), name)

					fmt.Printf("\nAvailable flags:\n")
					set.PrintDefaults()
				}
			}
		}
	}
}

// Name return the name of this subcommand.
func (h *Help) Name() string {
	return "help"
}

// Synopsis returns a one-line description of this command
func (h *Help) Synopsis() string {
	return "Show usage information."
}