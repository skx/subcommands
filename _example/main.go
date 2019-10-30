// This is an example which demonstrates using the subcommand-library.
//
// Usage is something like this:
//
//  _example help
//  _example help run
//  _example help version
//  _example run 1 2 3
//  _example version
//
package main

import (
	"flag"
	"fmt"

	"github.com/skx/subcommands"
)

//
// We're going to write an application with two sub-commands:
//
//  version -> Show our version
//  run     -> "Do stuff"
//
// To do that we'll defined a pair of objects, and implement
// a series of methods on each one.
//

//
// RUN
//

// runCommand is a subcommand that takes a single optional flag.
type runCommand struct {
	verbose bool
}

// Arguments adds per-command args to the object.
func (r *runCommand) Arguments(f *flag.FlagSet) {
	f.BoolVar(&r.verbose, "verbose", false, "Should we be verbose")

}

// Name returns the name of this subcommand.
func (r *runCommand) Name() string {
	return "run"
}

// Synopsis returns a one-line description of this command
func (r *runCommand) Synopsis() string {
	return "Runs some magic, and dumps its arguments."
}

// Execute is invoked if the user specifies `run` as the subcommand.
func (r *runCommand) Execute(args []string) {

	fmt.Printf("I am a running application!\n")

	fmt.Printf("Verbose flag is %t\n", r.verbose)

	for i, s := range args {
		fmt.Printf("Argument %d is %s\n", i, s)
	}
}

//
// VERSION
//

// versionCommand is a subcommand that takes no flags.
type versionCommand struct {
	subcommands.NoFlags
}

// Name returns the name of this subcommand.
func (h *versionCommand) Name() string {
	return "version"
}

// Synopsis returns a one-line description of this command
func (h *versionCommand) Synopsis() string {
	return "Show the application version."
}

// Execute is invoked if the user specifies `version` as the subcommand.
func (h *versionCommand) Execute(args []string) {
	fmt.Printf("I am application version 1.0\n")
}

//
// MAIN
//

//
// Register the subcommands, and run the one the user chose.
//
func main() {

	//
	// Register each of our subcommands.
	//
	subcommands.Register(&runCommand{})
	subcommands.Register(&versionCommand{})

	//
	// Execute the one the user chose.
	//
	subcommands.Execute()
}
