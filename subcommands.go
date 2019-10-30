// Package subcommands allows you to easily implement a simple CLI
// with well-defined and distinct subcommands.
//
// The implementation is pretty naive, but it is sufficient to
// implement a simple command.
package subcommands

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

// Subcommand is the interface which subcommands must implement.
//
// In brief a sub-command has a name, a function to invoke it, and
// the ability to define command-line flags which are specific to it.
//
// There is also a synopsis which can be displayed when a help
// command is used.
type Subcommand interface {

	// Name returns the name of this subcommand.
	Name() string

	// Synopsis returns a brief line of text.
	Synopsis() string

	// Arguments sets up any required arguments.
	Arguments(f *flag.FlagSet)

	// The function invoked if this command is invoked.
	Execute(args []string)
}

// NoFlags is a helper method which allows you to define sub-commands
// which take no flags.
//
// You still need to define `Name()`, `Synopsis()` and `Execute()`, but
// this saves a little needless typing.
type NoFlags struct {
}

// Arguments is a stub-method which registers no arguments.
func (nf *NoFlags) Arguments(flags *flag.FlagSet) {
}

// known stores the known-commands which have been registered.
var known []Subcommand

// init registers the built-in subcommands.
//
// We register `help` as a command by default.
func init() {
	Register(&Help{})
}

// Register adds a new subcommand to those which are available.
func Register(cmd Subcommand) {
	known = append(known, cmd)
}

// dump outputs a sorted list of all the known sub-commands, along
// with their brief synopsis.
func dump() {

	// Build up the list of names here.
	var names []string

	// store the usage here
	info := make(map[string]string)

	// Process each known sub-command
	for _, c := range known {
		names = append(names, c.Name())
		info[c.Name()] = c.Synopsis()
	}

	// Now sort the names
	sort.Strings(names)

	// Finally output them
	for _, name := range names {
		fmt.Printf("\t%s\t%s\n", name, info[name])
	}
}

// Execute launches the subcommand specified by the user.
func Execute() {

	//
	// Ensure the user specified a subcommand.
	//
	if len(os.Args) < 2 {
		fmt.Printf("You must specify the sub-command to execute:\n\n")
		dump()
		os.Exit(1)
	}

	//
	// Keep track of flags on a per-subcommand basis.
	//
	subcommandFlags := make(map[string]*flag.FlagSet)

	//
	// For each known sub-command
	//
	for _, c := range known {

		//
		// Get the name of the subcommand.
		//
		name := c.Name()

		//
		// Create a new flagset using that name.
		//
		set := flag.NewFlagSet(name, flag.ExitOnError)

		//
		// Setup the arguments, via the user-defined method.
		//
		c.Arguments(set)

		//
		// Now store the flagset away.
		//
		subcommandFlags[name] = set
	}

	//
	// Get the flags for the command the user chose.
	//
	subCmd, ok := subcommandFlags[os.Args[1]]
	if !ok {

		//
		// The user specified a subcommand which doesn't exist.
		//
		fmt.Printf("Invalid subcommand, available choices are:\n\n")
		dump()
		os.Exit(1)
	}

	//
	// Parse the flags setup by the user.
	//
	if err := subCmd.Parse(os.Args[2:]); err != nil {
		fmt.Printf("Error parsing flags %s\n", err.Error())
		os.Exit(1)
	}

	//
	// Execute the actual subcommand.
	//
	for _, c := range known {

		//
		// Use the name to store the flags in a hash.
		//
		if os.Args[1] == c.Name() {
			c.Execute(subCmd.Args())
			return
		}
	}
}
