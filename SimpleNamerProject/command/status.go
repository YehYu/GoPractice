package command

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli" // imports as package "cli"
	"os"
)

var (
	StatusCommand = cli.Command{
		Name:  "status",
		Usage: "狀態",
		Action: func(c *cli.Context) error {
			return getNames(c)
		},
	}
)

func getNames(c *cli.Context) error {
	os.Chdir(".")
	rc, _ := ioutil.ReadFile(c.GlobalString("namefile"))
	fmt.Printf("%s", rc)
	return nil
}
