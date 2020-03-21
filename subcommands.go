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
	"path"
	"sort"
)

// Subcommand is the interface which subcommands must implement.
//
// In brief a sub-command has a name, a function to invoke it, and
// the ability to define command-line flags which are specific to it.
//
type Subcommand interface {

	// Arguments sets up any required arguments.
	Arguments(f *flag.FlagSet)

	// Info is designed to returns the name, and brief
	// description of the command.
	Info() (string, string)

	// The function is invoked if this subcommand is invoked.
	//
	// The arguments are any non-flag arguments passed to the
	// subcommand, and the return value can be used as your
	// exit-code.
	Execute(args []string) int
}

// NoFlags is a helper method which allows you to define sub-commands
// which take no flags.
//
// You still need to define `Info`, and `Execute()`, but this saves a
// little needless typing.
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
	//
	// We can sort this list later.
	var names []string

	//
	// Store the synopsis for the command
	// in a map, so we can display it alongside
	// the name.
	//
	info := make(map[string]string)

	// Process each known sub-command
	for _, c := range known {

		// Get the name & synopsis
		name, synopsis := c.Info()

		// Save the name & info
		names = append(names, name)
		info[name] = synopsis
	}

	// Now sort the names
	sort.Strings(names)

	// Finally output each command in sorted-order
	// and the corresponding synopsis.
	for _, name := range names {
		fmt.Printf("\t%s\t%s\n", name, info[name])
	}
}

// Execute launches the subcommand specified by the user.
func Execute() int {

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
		name, _ := c.Info()

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
	// The subcommand can be specified via the name of the binary
	// or the first argument.
	//
	var valid []string
	valid = append(valid, path.Base(os.Args[0]))
	valid = append(valid, os.Args[1])

	//
	// Get the flags for the command the user chose.
	//
	var subCmd *flag.FlagSet

	//
	// The argument offset
	//
	var args int

	//
	// The sub-command name
	//
	var subCmdName string

	//
	// Try to match either attempt.
	//
	for i, attempt := range valid {

		var ok bool
		subCmd, ok = subcommandFlags[attempt]
		if ok {
			args = i + 1
			subCmdName = attempt
		}
	}

	//
	// If we didn't find the args then we didn't find a subcommand
	//
	if args == 0 {
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
	if err := subCmd.Parse(os.Args[args:]); err != nil {
		fmt.Printf("Error parsing flags %s\n", err.Error())
		os.Exit(1)
	}

	//
	// Execute the actual subcommand.
	//
	for _, c := range known {

		//
		// Get the name of the available sub-command.
		//
		name, _ := c.Info()

		//
		// If it is what the user wanted, then invoke it.
		//
		if subCmdName == name {
			return (c.Execute(subCmd.Args()))
		}
	}

	return 0
}
