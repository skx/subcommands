[![GoDoc](https://godoc.org/github.com/skx/subcommands?status.svg)](http://godoc.org/github.com/skx/subcommands)
[![Go Report Card](https://goreportcard.com/badge/github.com/skx/subcommands)](https://goreportcard.com/report/github.com/skx/subcommands)
[![license](https://img.shields.io/github/license/skx/subcommands.svg)](https://github.com/skx/subcommands/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/subcommands.svg)](https://github.com/skx/subcommands/releases/latest)


# subcommands

This is a simple package which allows you to write a CLI with distinct subcommands.


## Overview

Using this library you can enable your command-line application to have a number of subcommands, allowing things like this to be executed:

    $ application one
    $ application two [args]
    $ application help

In addition to allowing the user to specify a sub-command via the first argument it will also allow a default to be used if your binary has the same name as a sub-command.

For example if you had a binary named `gobox` you could create a symlink called `ls`:

    $ ln -s gobox ls
    $ ./ls

Here running `ls` is the same as running `gobox ls`, and argument parsing would work the same too:

    $ gobox ls --foo
    $ ./ls --foo


## Bash Completion

All applications using this library will find it easy to generate a bash-completion script, via the following addition to their init-file:

    $ source <(application bash-completion)

The generated completion-script will allow TAB-completion of the sub-commands, as well as their options.



## Rationale

There are several frameworks for building a simple CLI application, such as Corba.  But those are relatively heavyweight.

This is designed to implement the minimum required support:

* Allow an application to register sub-commands:
  * `foo help`
    * Show help
  * `foo version`
    * Show the version
  * `foo server`
    * Launch a HTTP-server
  * etc.
* Allow flags to be defined on a per-subcommand basis.
* Allow the sub-commands to process any remaining arguments.
* Allow trivial bash-completion scripts to be generated.


## Example

There is a simple example defined in [_example/main.go](_example/main.go).
