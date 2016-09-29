package lib

import (
	"fmt"
	"github.com/urfave/cli"
)

// Start ...
func Start(c *cli.Context) {

	var (
		commit, push int
		group        string
	)

	fmt.Println("This session will have the following settings:")

	// if commit set value
	if c.IsSet("commit") && c.Int("commit") != 0 {
		fmt.Println("  - Commit:", c.Int("commit"))
		commit = c.Int("commit")
	} else if c.IsSet("commit") == false {
		commit = -1
	} else {
		fmt.Println("  - Default Push: 5")
		commit = 10
	}

	// if push set value
	if c.IsSet("push") && c.Int("push") != 0 {
		fmt.Println("  - Push:", c.Int("push"))
		push = c.Int("push")
	} else if c.IsSet("push") == false {
		push = -1
	} else {
		fmt.Println("  - Default Push: 60")
		push = 60
	}

	// If group exists grab group name and set the values
	if c.Args().First() != "" {
		fmt.Println("  - Group:", c.Args().First())
		group = c.Args().First()
		// settings set on
		commit = 2
		push = 2
	} else {
		fmt.Println("  - Group Not Set")
		group = "-1"
	}

	fmt.Println(" commit ", commit)
	fmt.Println(" push ", push)
	fmt.Println(" group ", group)

	// Call to API

	/*
	   Check update times
	     - Local vs API
	*/

	// Check if .forgitConf.json exists

	// Read .forgitConf.json file

	// Grab all repo names from config with status 1

	// Set Forgit path

	// Check if Forgit path is valid

	/*
	   Go to each repos
	   - Check if .git
	   - git status short
	     - to slice
	   - file read status files into map
	*/

}
