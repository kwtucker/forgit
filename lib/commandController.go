package lib

import (
	"github.com/kwtucker/fileReader"
	"log"
	"os"
	osuser "os/user"
	"strings"
	"sync"
	"time"
)

//CommandController dispatches the commands
func CommandController(settingObj Setting, path string, repos []SettingRepo, uuid string, gitCommand string) {
	// Current home dir of user OS
	homeDir, err := osuser.Current()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {
		// Read and update the config file
		FileExist(homeDir.HomeDir+"/.forgitConf.json", path, homeDir.HomeDir, uuid, "no")

		for r := range repos {

			var (
				err error
				wg  sync.WaitGroup
			)

			// Go to the current repo directory and get the current branch
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
				ctime, err := GetCurrentCPTimeMin(settingObj, "commit")
				if err != nil {
					log.Println(err)
				}
				if ctime != 0 {
					// a delay in the for loop
					Ticker(ctime)
				} else {
					Ticker(settingObj.SettingAddPullCommit.TimeMin)
				}

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
			case "push":
				ptime, err := GetCurrentCPTimeMin(settingObj, "push")
				if err != nil {
					log.Println(err)
				}
				if ptime != 0 {
					// a delay in the for loop
					Ticker(ptime)
				} else {
					Ticker(settingObj.SettingPush.TimeMin)
				}
				// a delay in the for loop
				// Ticker(settingObj.SettingPush.TimeMin)
				wg.Add(1)
				go GitPushPull(path+repos[r].Name, branchName, "push", &wg)
			}
			wg.Wait()
		}
	}

}
