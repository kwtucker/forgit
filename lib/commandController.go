package lib

import (
	"encoding/json"
	"github.com/kwtucker/fileReader"
	"io/ioutil"
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
		var (
			dataUser []User
			setdata  []Setting
			repoArr  []SettingRepo
		)
		configfile, err := ioutil.ReadFile(homeDir.HomeDir + "/.forgitConf.json")
		if err != nil {
			os.Exit(1)
		}
		json.Unmarshal(configfile, &dataUser)

		// Grab most recent data and set it to the datauser
		curldata, err := Curlforgit("no", dataUser[0].ForgitID)
		if err != nil {
			log.Println(err)
		}
		if len(curldata) > 200 {
			// Format curl data and set it to settings array
			err = json.Unmarshal(curldata, &setdata)
			for set := range setdata {
				if setdata[set].Name == settingObj.Name {
					settingObj = setdata[set]
					log.Println(len(settingObj.Repos))
					for i := range settingObj.Repos {
						if settingObj.Repos[i].Status == 1 {
							repoArr = append(repoArr, settingObj.Repos[i])
						}
					}
					if len(repoArr) == 0 {
						log.Println(": You don't have any repos to automate.\nOr you don't have any selected in setting group.")
						os.Exit(1)
					}
					break
				}
			}
		}

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
