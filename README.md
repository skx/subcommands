# subcommands

This is a simple package which allows you to write a CLI with git-like subcommands.

## Overview


## Rationale

There are several frameworks for building a simple CLI application, such as
Corba.  But those are relatively heavyweight.

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

I think this is a clean implementation, however I appreciate we don't
yet have any tests.

## Example

There is a simple example defined in [_example/main.go](_example/main.go).
