package command

import (
	"fmt"
	"os"
	"strconv"

	"github.com/YehYu/GoPractice/SimpleNamerProject/provider"
	"github.com/urfave/cli" // imports as package "cli"
)

var (
	GenerateCommand = cli.Command{
		Name:  "generate",
		Usage: "產生假名",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "num",
				Value: "10",
				Usage: "產生數量",
			},
		},
		Action: func(c *cli.Context) error {
			return generate(c)

		},
	}
)

func generate(c *cli.Context) error {
	num, err := strconv.Atoi(c.String("num"))

	if err != nil {
		return err
	}

	fmt.Println("Generate " + strconv.Itoa(num))
	os.Chdir(".")
	ns, _ := provider.GetNames(c.GlobalString("namefile"))
	generator := provider.Create()
	generator.Names = ns
	for i := 0; i < num; i++ {
		fmt.Println(generator.Name())
	}

	return nil
}
