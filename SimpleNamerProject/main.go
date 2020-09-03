package main

import (
	"os"

	"github.com/YehYu/GoPractice/SimpleNamerProject/command"
	"github.com/urfave/cli" // imports as package "cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Namer"
	app.Commands = []cli.Command{
		command.GenerateCommand,
		command.StatusCommand,
	}

	app.Run(os.Args)
}
