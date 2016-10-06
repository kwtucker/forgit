package lib

import (
	"fmt"
	"log"
	"os"
	"sync"
)

//CommandController dispatches the commands
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

			_, err = Status(path + repos[r].Name)
			if err != nil {
				log.Println(err)
			}
			var wg sync.WaitGroup
			wg.Add(1)

			switch gitCommand {
			case "commit":
				fmt.Println("commit")
			case "push":
				fmt.Println("push")
				go GitPushPull(path+repos[r].Name, branchName, "push", &wg)
			}
			wg.Wait()
		}
	}

}
