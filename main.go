package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.Usage = Usage

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.Before = VerboseGlobalOptionHandler

	app.Run(os.Args)
}
