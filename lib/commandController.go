package lib

import (
	"encoding/json"
	"fmt"
	"github.com/kwtucker/fileReader"
	"io/ioutil"
	"log"
	"os"
	osuser "os/user"
	"strconv"
	"strings"
	"sync"
	"time"
)

//CommandController dispatches the commands
func CommandController(settingObj Setting, path string, repos []SettingRepo, uuid string, internet bool, gitCommand string) {
	// Current home dir of user OS
	homeDir, err := osuser.Current()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var commitCounter, pushCounter int

	for {
		var (
			dataUser []User
			setdata  []Setting
			repoArr  []SettingRepo
		)

		// Read config
		configfile, err := ioutil.ReadFile(homeDir.HomeDir + "/.forgitConf.json")
		if err != nil {
			os.Exit(1)
		}
		json.Unmarshal(configfile, &dataUser)

		// If the internet is true it will call to the server and update user object.
		if internet {
			// Grab most recent data and set it to the datauser
			curldata, err := Curlforgit("no", dataUser[0].ForgitID)
			if err != nil {
				log.Println(err)
			}

			// if it is bad UUID was entered. The curl data will always be 42 bytes for bad result.
			if len(curldata) == 42 {
				var aerr APIError
				err = json.Unmarshal(curldata, &aerr)
				if err != nil {
					log.Println(err)
				}
				// If the forgit Id is wrong
				if aerr.Status == 401 {
					fmt.Println("Bad UUID credentials,")
					fmt.Println(" 1. Try forgit init again and make sure to copy all the UUID from the dashboard on the browser.")
					fmt.Println(" 2. If you did not get the CLI you are using from forgit.whalebyte.com, be sure to \nlogin to forgit.whalebyte.com and get your own UUID from the dashboard.")
					if settingObj.OnError == 1 {
						m := &Message{
							Title: "User UUID Wrong",
							Body:  "Refer to your Terminal",
						}
						Notify(*m)
					}
					os.Exit(1)
				}
			}

			// if it is greater than 200 data was updated.
			if len(curldata) > 200 {
				// Format curl data and set it to settings array
				err = json.Unmarshal(curldata, &setdata)
				dn := time.Now().UTC().Unix()
				dateNow := strconv.FormatInt(dn, 10)
				dataUser[0].ForgitPath = path
				dataUser[0].Settings = setdata
				dataUser[0].UpdateTime = dateNow
				databytes, err := json.MarshalIndent(dataUser, "", "    ")
				if err != nil {
					log.Println(err.Error())
					os.Exit(1)
				}

				// Write to file with updated info
				err = ioutil.WriteFile(homeDir.HomeDir+"/.forgitConf.json", databytes, 0644)
				if err != nil {
					log.Println(err.Error())
					os.Exit(1)
				}
				for set := range setdata {
					if setdata[set].Name == settingObj.Name {
						settingObj = setdata[set]
						for i := range settingObj.Repos {
							if settingObj.Repos[i].Status == 1 {
								repoArr = append(repoArr, settingObj.Repos[i])
							}
						}
						if len(repoArr) == 0 {
							log.Println(": You don't have any repos to automate.\n" +
								"\tOr you don't have any selected in setting group.\n" +
								"\tSelect repos in the " + settingObj.Name + " workspace and restart. forgit start")
							if settingObj.OnError == 1 {
								m := &Message{
									Title: "Setting Repo Error",
									Body:  "No repos to automate",
								}
								Notify(*m)
							}
							os.Exit(1)
						}
						break
					}
				}
			}
		}

		// If the length of repos in the Forgit dir is 0 stop the app.
		if len(repos) == 0 {
			log.Println(": You don't have any repos to automate.\n" +
				"\tOr you don't have any selected in setting group.\n" +
				"\tSelect repos in the " + settingObj.Name + " workspace and restart. forgit start")
			os.Exit(1)
		}

		// Delays the first start commit and push
		if gitCommand == "commit" && commitCounter == 0 {
			Ticker(settingObj.SettingAddPullCommit.TimeMin)
			commitCounter++
		}
		if gitCommand == "push" && pushCounter == 0 {
			Ticker(settingObj.SettingPush.TimeMin)
			pushCounter++
		}

		// Loop over the repos in the Forgit directory
		for r := range repos {

			var (
				err       error
				wg        sync.WaitGroup
				dataSlice []string
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
			status, err := GitStatus(path + repos[r].Name)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(1 * time.Second)

			switch gitCommand {
			case "commit":
				if commitCounter >= 1 {
					// Get current times for commit.
					ctime, noteerr, notecommit, _, err := GetCurrentCommitPushTimeMin(settingObj, "commit")
					if err != nil {
						log.Println(err)
					}

					// Only if there is internet will it pull.
					if internet {
						wg.Add(1)
						go GitPushPull(path+repos[r].Name, branchName, "pull", &wg, 0, noteerr)
						time.Sleep(4 * time.Second)
					}
					if settingObj.OnCommit == 1 {
						m := &Message{
							Title: "Save Files",
							Body:  "Forgit Event In 15 seconds",
						}
						Notify(*m)
						time.Sleep(15 * time.Second)
					}

					// Read file, git add, git commit each file in the git status output.
					for _, s := range status {
						// reads the file it is currently on.
						dataSlice = fileReader.ReadFile(path+repos[r].Name+"/"+s, false)
						formatSlice := strings.Join(dataSlice, "\n-")
						wg.Add(2)
						go GitAdd(s, &wg)
						time.Sleep(500 * time.Millisecond)
						go GitCommit(formatSlice, &wg, notecommit, noteerr)
						time.Sleep(500 * time.Millisecond)
					}

					if ctime != 0 {
						Ticker(ctime)
					} else {
						Ticker(settingObj.SettingAddPullCommit.TimeMin)
					}
				}
			case "push":
				if pushCounter >= 1 {
					// Get current push time.
					ptime, noteerr, _, notepush, err := GetCurrentCommitPushTimeMin(settingObj, "push")
					if err != nil {
						log.Println(err)
					}
					wg.Add(1)
					go GitPushPull(path+repos[r].Name, branchName, "push", &wg, notepush, noteerr)
					if ptime != 0 {
						// a delay in the for loop
						Ticker(ptime)
					} else {
						Ticker(settingObj.SettingPush.TimeMin)
					}
				}
			}
			wg.Wait()
		}
	}
}
