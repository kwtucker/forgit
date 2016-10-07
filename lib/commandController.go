package lib

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

//CommandController dispatches the commands
func CommandController(gtime int, path string, repos []SettingRepo, gitCommand string) {

	for {
		// a delay in the for loop
		Ticker(gtime)
		// Where the push code is going
		for r := range repos {
			var (
				err error
				wg  sync.WaitGroup
			)

			err = os.Chdir(path + repos[r].Name)
			branchName, err := GetCurrentBranch(path + repos[r].Name)
			if err != nil {
				log.Println(err)
			}

			// get status slice
			// Wait tell the status slice is generated
			status, err := Status(path + repos[r].Name)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(1 * time.Second)

			switch gitCommand {
			case "commit":
				wg.Add(1)
				go GitPushPull(path+repos[r].Name, branchName, "pull", &wg)
				time.Sleep(2 * time.Second)
				for _, s := range status {
					wg.Add(1)
					go GitAdd(s, &wg)
				}
				fmt.Println("commit")
			case "push":
				wg.Add(1)
				go GitPushPull(path+repos[r].Name, branchName, "push", &wg)
			}
			wg.Wait()
		}
	}

}
