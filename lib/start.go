package lib

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	osuser "os/user"
)

func settingGroupsCheck(groupName string, dataUser User) (Setting, bool) {
	for u := range dataUser.Settings {
		if dataUser.Settings[u].Name == groupName {
			return dataUser.Settings[u], true
		}
	}
	return dataUser.Settings[0], false
}

// Start ...
func Start(c *cli.Context) {

	var (
		commit, push int
		group        string
		dataUser     []User
	)
	// Check if .forgitConf.json exists
	homeDir, err := osuser.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err = os.Stat(homeDir.HomeDir + "/.forgitConf.json"); os.IsNotExist(err) {
		fmt.Println()
		fmt.Println("* Haven't started yet!")
		fmt.Println("* Please first run --> fgt init ")
		fmt.Println()
		return
	}

	// Read .forgitConf.json file
	configfile, err := ioutil.ReadFile(homeDir.HomeDir + "/.forgitConf.json")
	if err != nil {
		os.Exit(1)
	}

	// data from api
	json.Unmarshal(configfile, &dataUser)

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

		settingObj, setExist := settingGroupsCheck(group, dataUser[0])
		if !setExist {
			fmt.Println()
			fmt.Println("* Did Not Start!")
			fmt.Println("* Setting Workspace group does not exist.")
			fmt.Println()
			return
		}
		// set group vars
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

	// git byte array from MarshalIndent
	configDataBytes, err := json.MarshalIndent(dataUser, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(dataUser[0].GithubID)
	fmt.Println(string(configDataBytes))

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
