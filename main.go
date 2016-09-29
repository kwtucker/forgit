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
	app.Name = "fgt"
	app.Usage = "fgt"
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
			Name:        "init",
			Aliases:     []string{"i"},
			Usage:       "fgt init",
			Description: "Creates the config file in your home directory that the app uses.",
			Action: func(c *cli.Context) error {
				lib.Init()
				return nil
			},
		},
		{
			Name:        "start",
			Aliases:     []string{"s"},
			Usage:       "fgt start",
			Description: "Starts app and automates based on you forgit settings.",
			ArgsUsage:   "NUMBER-MINUTES",
			Subcommands: []cli.Command{
				{
					Name:        "group",
					Aliases:     []string{"g"},
					Usage:       "fgt start group GROUP-NAME",
					Description: "Set Workspace setting group",
					Action: func(c *cli.Context) error {
						lib.Start(c)
						return nil
					},
				},
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "commit, c",
					Value: "5",
					Usage: "--> Set commit repeat time MINUTES |",
				},
				cli.StringFlag{
					Name:  "push, p",
					Value: "60",
					Usage: "--> Set push repeat time MINUTES   |",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("To select a Workspace")
				fmt.Println("fgt start group GROUP-NAME", "--> That will select one of your settings groups")
				lib.Start(c)
				return nil
			},
		},
		{
			Name:    "sp",
			Aliases: []string{"sp"},
			Usage:   "Stop Forgit",
			Action: func(c *cli.Context) error {
				fmt.Println("fgt stop is Coming Soon")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
