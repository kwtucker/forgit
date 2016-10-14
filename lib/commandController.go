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

		// If the length of repos in the Forgit dir is 0 stop the app.
		if len(repos) == 0 {
			log.Println(": You don't have any repos to automate.\nOr you don't have any selected in setting group.")
			os.Exit(1)
		}

		// Loop over the repos in the Forgit directory
		for r := range repos {

			var (
				err error
				wg  sync.WaitGroup
			)

			// Go to the current repo directory and get the current branch
			err = os.Chdir(path + repos[r].Name)
			if err != nil {
				log.Println(": Repo:", repos[r].Name, "does not exist!")
				os.Exit(1)
			}

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
				ctime, noteerr, notecommit, _, err := GetCurrentCPTimeMin(settingObj, "commit")
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
				go GitPushPull(path+repos[r].Name, branchName, "pull", &wg, 0, noteerr)
				time.Sleep(4 * time.Second)

				for _, s := range status {
					// reads the file it is currently on. Takes 15 seconds
					dataSlice := fileReader.ReadFile(path + repos[r].Name + "/" + s)
					formatSlice := strings.Join(dataSlice, "\n-")
					wg.Add(2)
					go GitAdd(s, &wg)
					time.Sleep(500 * time.Millisecond)
					go GitCommit(formatSlice, &wg, notecommit, noteerr)
				}
			case "push":
				ptime, noteerr, _, notepush, err := GetCurrentCPTimeMin(settingObj, "push")
				if err != nil {
					log.Println(err)
				}
				if ptime != 0 {
					// a delay in the for loop
					Ticker(ptime)
				} else {
					Ticker(settingObj.SettingPush.TimeMin)
				}

				wg.Add(1)
				go GitPushPull(path+repos[r].Name, branchName, "push", &wg, notepush, noteerr)
				time.Sleep(4 * time.Second)
			}
			wg.Wait()
		}
	}

}
