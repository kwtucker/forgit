package main

import (
	"fmt"
	"github.com/kwtucker/forgit/lib"
	"github.com/urfave/cli"
	"os"
)

func main() {
	logo :=
		`
   ______                  _  _
   |  ___|                (_)| |
   | |_  ___   _ __  __ _  _ | |_
   |  _|/ _ \ |  __|/ _  || || __|
   | | | (_) || |  | (_| || || |_
   \_|  \___/ |_|   \__, ||_| \__|
                     __/ |
                    |___/
`

	app := cli.NewApp()
	app.Name = "forgit"
	app.Author = "Kevin Tucker\n\t https://github.com/kwtucker"
	app.Usage = "forgit"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		fmt.Println(logo)
		fmt.Println("\tWelcome to Forgit CLI")
		fmt.Println("\t  Let's GIT to it!")
		fmt.Println("     > forgit help [-h, h, --help]")
		fmt.Println()
		return nil
	}

	// Commands that the user can run. Using the github.com/urfave/cli.
	app.Commands = []cli.Command{
		{
			Name:        "init",
			Aliases:     []string{"i"},
			Usage:       "forgit init",
			Description: "Creates the config file in your home directory that the app uses.",
			Action: func(c *cli.Context) error {
				lib.Init()
				return nil
			},
		},
		{
			Name:        "start",
			Aliases:     []string{"s"},
			Usage:       "forgit start",
			Description: "Starts app and automates based on you forgit settings.",
			ArgsUsage:   "NUMBER-MINUTES",
			Subcommands: []cli.Command{
				{
					Name:        "group",
					Aliases:     []string{"g"},
					Usage:       "forgit start group GROUP-NAME",
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
				fmt.Println("-->  forgit start group GROUP-NAME")
				lib.Start(c)
				return nil
			},
		},
		{
			Name:    "stop",
			Aliases: []string{"sp"},
			Usage:   "To stop the app you must do ONE of the following: \n\t\t1. Close the forgit shell window.\n\t\t2. Control-c in the forgit window.",
			Action: func(c *cli.Context) error {
				fmt.Println("To stop the app you must do ONE of the following: \n\t1. Close the forgit shell window.\n\t2. Control-c in the forgit window.")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
