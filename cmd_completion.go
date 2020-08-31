// cmd_completion.go: Helpers for completion

package subcommands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CommandList is a structure which implements the built-in "commands" subcommand
type CommandList struct {

	// We embed the NoFlags option, because we accept no
	// command-line flags.
	NoFlags
}

// Execute the command.
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
func (c *CommandList) Info() (string, string) {
	return "commands", "Show all available sub-commands."
}

// BashCompletion is a structure which implements the built-in
// "bash-completion" subcommand.
type BashCompletion struct {

	// We embed the NoFlags option, because we accept no
	// command-line flags.
	NoFlags
}

// Execute the command.
func (bc *BashCompletion) Execute(args []string) int {

	tmpl := `
_subcommands_#Command#()
{
    local cur
    COMPREPLY=()

    # Variable to hold the current word
    cur="${COMP_WORDS[COMP_CWORD]}"

    # The first argument is one of the available sub-commands.
    if [ $COMP_CWORD = 1 ]; then

        local subs=$(#Command# commands)
        COMPREPLY=($(compgen -W "${subs}" $cur))
    else

        # If we see a dash complete from the available flags,
        # otherwise a file/directory.
        if [[ "$cur" =~ ^-.* ]];  then
            local flags="$(evalfilter help ${COMP_WORDS[1]} | awk '{print $1}' | grep -- -)"
            COMPREPLY=($(compgen -W "${flags}" -- "$cur"))
        else
            COMPREPLY=($(compgen -f -- ${cur}))
        fi
   fi
}

complete -F _subcommands_#Command# #Command#
`

	output := strings.ReplaceAll(tmpl, "#Command#", filepath.Base(os.Args[0]))
	fmt.Printf("%s\n", output)
	return 0
}

// Info returns the name of this subcommand, along with a one-line synopsis.
func (bc *BashCompletion) Info() (string, string) {
	return "bash-completion", "Generate and output a bash completion-script."
}
