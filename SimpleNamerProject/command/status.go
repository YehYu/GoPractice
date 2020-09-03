package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli" // imports as package "cli"
)

var (
	StatusCommand = cli.Command{
		Name:  "status",
		Usage: "狀態",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "namefile",
				Value: "name.yml",
				Usage: "名字庫",
			},
		},
		Action: func(c *cli.Context) error {
			return getNames(c)
		},
	}
)

func getNames(c *cli.Context) error {
	os.Chdir(".")
	rc, _ := ioutil.ReadFile(c.String("namefile"))
	fmt.Printf("%s", rc)
	return nil
}
