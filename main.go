package main

import (
	"fmt"
	"github.com/kwtucker/fgt/lib"
	"github.com/urfave/cli"
	"os"
)

func main() {
	logo :=
		`
    '.:/:.'
    ooo+ooo
    ossssso
  '.'-/o/-'.'
//+++/- -/+++//
sssosso ossosss
/sssso: :ossss/
  .-'.:/:.'-.
    oo+++oo
    ossssso
    ':+s+:'
`

	app := cli.NewApp()
	app.Name = "Forgit CLI"
	app.Usage = "Never Forget To Commit"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		fmt.Println(logo)

		fmt.Println("\tWelcome to Forgit CLI 🍺")
		fmt.Println("\t  Let's GIT to it!")
		fmt.Println("     > fgt help [-h, h, --help]")
		fmt.Println()
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name: "init",
			Action: func(c *cli.Context) error {
				lib.Init()
				return nil
			},
		},
		{
			Name:        "start",
			Aliases:     []string{"st"},
			Usage:       "Start Forgit",
			Description: "Starts app and automates based on you forgit settings.",
			ArgsUsage:   "NUMBER-MINUTES",
			Action: func(c *cli.Context) error {
				fmt.Println("fgt start is Coming Soon")
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "commit, c",
					Value: "5",
					Usage: "--> Set commit repeat time MINUTES |",
				},
				cli.StringFlag{
					Name:  "push, p",
					Value: "5",
					Usage: "--> Set Push repeat time *minutes |",
				},
			},
		},
		{
			Name:    "stop",
			Aliases: []string{"sp"},
			Usage:   "Stop Forgit",
			Action: func(c *cli.Context) error {
				// fmt.Println("completed task: ", c.Args().First())
				fmt.Println("fgt stop is Coming Soon")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
