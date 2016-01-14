package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/griffin-stewie/ta/command"
)

// GlobalFlags is
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "backlog-token",
		Usage:  "API Token to access Backlog API",
		EnvVar: "BACKLOG_TOKEN",
	},
	cli.StringFlag{
		Name:   "backlog-endpoint",
		Usage:  "API endpoint to access Backlog API",
		EnvVar: "BACKLOG_API",
	},
	cli.StringFlag{
		Name:   "chatwork-token",
		Usage:  "API Token to access ChatWork API",
		EnvVar: "CHATWORK_API_TOKEN",
	},
	cli.BoolFlag{
		Name:  "verbose",
		Usage: "Show debug logs",
	},
}

// Commands is
var Commands = []cli.Command{
	{
		Name:   "list",
		Usage:  "",
		Action: command.CmdList,
		Flags:  []cli.Flag{},
	},
}

// CommandNotFound ... command not found
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

// VerboseGlobalOptionHandler runs before subcommands
func VerboseGlobalOptionHandler(c *cli.Context) error {
	if c.GlobalBool("verbose") {
		SetupLogger(Debug)
	} else {
		SetupLogger(Error)
	}

	return nil
}
