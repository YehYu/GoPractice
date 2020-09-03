package main

import (
	"os"

	"github.com/YehYu/GoPractice/SimpleNamerProject/command"
	"github.com/urfave/cli" // imports as package "cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Namer"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "namefile, n",
			Value: "name.yml",
			Usage: "名字庫",
		},
	}
	app.Commands = []cli.Command{
		command.GenerateCommand,
		command.StatusCommand,
	}

	app.Run(os.Args)
}
