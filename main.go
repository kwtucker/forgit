package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func hello() {
	fmt.Println("hello")
}

func main() {
	app := cli.NewApp()
	app.Name = "Forgit CLI"
	app.Usage = "Never Forget To Commit"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Welcome, Let's GIT to it.")
		return nil
	}

	app.Commands = []cli.Command{

		{
			Name:    "stop",
			Aliases: []string{"sp"},
			Usage:   "Stop Forgit",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				hello()
				return nil
			},
		},
		{
			Name:        "start",
			Aliases:     []string{"st"},
			Usage:       "Start Forgit",
			Description: "Starts app and automates based on you forgit settings.",
			ArgsUsage:   "NUMBER-MINUTES",
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
	}
	app.Run(os.Args)
}
