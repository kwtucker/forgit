package lib

import (
	"fmt"
	"log"
	"os"
)

//GitPush is simply a git push command
func CommandController(time int, path string, repos []SettingRepo, gitCommand string) {

	for {
		// a delay in the for loop
		Ticker(time)
		// Where the push code is going
		for r := range repos {
			var err error
			err = os.Chdir(path + repos[r].Name)
			branchName, err := GetCurrentBranch(path + repos[r].Name)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(branchName)

			status, err := Status(path + repos[r].Name)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(status)

			switch gitCommand {
			case "commit":
				fmt.Println("commit")
			case "push":
				fmt.Println("push")
			}

		}
	}

}
