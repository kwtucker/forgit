package lib

import (
	"fmt"
	"github.com/kwtucker/fileReader"
	"log"
	"os"
	osuser "os/user"
	"strings"
	"sync"
	"time"
)

//CommandController dispatches the commands
func CommandController(settingObj Setting, path string, repos []SettingRepo, gitCommand string) {
	homeDir, err := osuser.Current()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for {
		FileExist(homeDir.HomeDir+"/.forgitConf.json", path, homeDir.HomeDir)

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
				commitTime, err := GetCurrentCPTimeMin(settingObj, "commit")
				if err != nil {
					log.Println(err)
				}
				fmt.Println(commitTime)
				// a delay in the for loop
				Ticker(commitTime)
				wg.Add(1)
				go GitPushPull(path+repos[r].Name, branchName, "pull", &wg)
				time.Sleep(4 * time.Second)
				for _, s := range status {
					// reads the file it is currently on. Takes 15 seconds
					dataSlice := fileReader.ReadFile(path + repos[r].Name + "/" + s)
					formatSlice := strings.Join(dataSlice, "\n-")
					wg.Add(2)
					go GitAdd(s, &wg)
					time.Sleep(500 * time.Millisecond)
					go GitCommit(formatSlice, &wg)
				}
				fmt.Println("commit")
			case "push":
				pushTime, err := GetCurrentCPTimeMin(settingObj, "push")
				if err != nil {
					log.Println(err)
				}
				fmt.Println(pushTime)
				// a delay in the for loop
				Ticker(pushTime)
				wg.Add(1)
				go GitPushPull(path+repos[r].Name, branchName, "push", &wg)
			}
			wg.Wait()
		}
	}

}
