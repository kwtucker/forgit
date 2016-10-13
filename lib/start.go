package lib

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	osuser "os/user"
	"sync"
)

func settingGroupsCheck(groupName string, dataUser User) (Setting, bool) {
	for u := range dataUser.Settings {
		if dataUser.Settings[u].Name == groupName {
			return dataUser.Settings[u], true
		}
	}
	return dataUser.Settings[0], false
}

// ForgitDirReposNames ...
func ForgitDirReposNames(path string) []string {
	var forgitDirReposNameSlice []string
	// read forgit directory
	repos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		if repo.IsDir() {
			files, err := ioutil.ReadDir(path + repo.Name())
			if err != nil {
				log.Fatal(err)
			}
			for _, f := range files {
				if f.Name() == ".git" {
					forgitDirReposNameSlice = append(forgitDirReposNameSlice, repo.Name())
				}
			}
		}
	}
	return forgitDirReposNameSlice
}

// InternetCheck Checks if internet exists
func InternetCheck() bool {
	_, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println(err)
		fmt.Println("No internet")
		return false
	}
	return true
}

// Start ...
func Start(c *cli.Context) {

	var (
		commit, push int
		group        string
		dataUser     []User
		settingObj   Setting
		settingRepos []SettingRepo
		settingRepo  SettingRepo
		setExist     bool
	)

	// Internet Check
	internetConnection := InternetCheck()

	// get Home directory path
	homeDir, err := osuser.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check if config file exists.
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

	// data from configfile
	json.Unmarshal(configfile, &dataUser)

	fmt.Println("This session will have the following settings:")

	// If group exists grab group name and set the values
	if c.Args().First() != "" {
		group = c.Args().First()
		settingObj, setExist = settingGroupsCheck(group, dataUser[0])
		if setExist == false {
			fmt.Println()
			fmt.Println("* Did Not Start!")
			fmt.Println("* Setting Workspace group does not exist.")
			fmt.Println()
			return
		}
	} else {
		group = "-1"
	}

	// if commit set value
	if c.IsSet("commit") && c.Int("commit") != 0 && group == "-1" {
		commit = c.Int("commit")
	} else {
		commit = -1
	}

	// if push set value
	if c.IsSet("push") && c.Int("push") != 0 && group == "-1" {
		push = c.Int("push")
	} else {
		push = -1
	}

	FileExist(homeDir.HomeDir+"/.forgitConf.json", dataUser[0].ForgitPath+"Forgit/", homeDir.HomeDir)

	/*
	   Check update times
	     - Local vs API
	*/

	// git byte array from MarshalIndent
	// configDataBytes, err := json.MarshalIndent(dataUser, "", "    ")
	_, err = json.MarshalIndent(dataUser, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// if no push,commit, or group set. Grab setting with status of 1
	if c.IsSet("push") == false && c.IsSet("commit") == false && group == "-1" {
		for s := range dataUser[0].Settings {
			if dataUser[0].Settings[s].Status == 1 {
				settingObj = dataUser[0].Settings[s]
			}
		}
	}

	// Send setting repos with status 1 to a slice.
	for r := range settingObj.Repos {
		if settingObj.Repos[r].Status == 1 {
			settingRepos = append(settingRepos, settingObj.Repos[r])
		}
	}

	// If push or commit set and no group build a struct.
	// This only lasts per session
	if settingObj.Name == "" {
		var (
			setPush   SettingPush
			setCommit SettingAddPullCommit
		)

		if push != -1 {
			setPush = SettingPush{
				TimeMin: push,
			}
		} else {
			setPush = SettingPush{
				TimeMin: push,
			}
		}

		if commit != -1 {
			setCommit = SettingAddPullCommit{
				TimeMin: commit,
			}
		} else {
			setCommit = SettingAddPullCommit{
				TimeMin: commit,
			}
		}

		// Set all the repos in the forgit directory to status 1 on.
		repoNames := ForgitDirReposNames(dataUser[0].ForgitPath + "/Forgit/")
		for r := range repoNames {
			settingRepo = SettingRepo{
				GithubRepoID: r,
				Name:         repoNames[r],
				Status:       1,
			}
			settingRepos = append(settingRepos, settingRepo)
		}

		// Struct build for the session setting
		settingObj = Setting{
			Name:   "fgtDefault",
			Status: 1,
			SettingNotifications: SettingNotifications{
				OnError:  1,
				OnCommit: 1,
				OnPush:   1,
			},
			SettingAddPullCommit: setCommit,
			SettingPush:          setPush,
			Repos:                settingRepos,
		}
	}

	fmt.Println("dataUser ID ->", dataUser[0].GithubID)
	// fmt.Println(string(configDataBytes))
	fmt.Println("settingObj Name ->", settingObj.Name)
	fmt.Println("settingObj commit time ->", settingObj.SettingAddPullCommit.TimeMin)
	fmt.Println("settingObj push time ->", settingObj.SettingPush.TimeMin)
	fmt.Println("internetConnection ->", internetConnection)

	// This will only pass in the repos that exist in the Forgit Directory and that are set to 1
	var automateRepos []SettingRepo
	repoArr := ForgitDirReposNames(dataUser[0].ForgitPath)
	for r := range settingObj.Repos {
		for s := range repoArr {
			if settingObj.Repos[r].Name == repoArr[s] && settingObj.Repos[r].Status == 1 {
				automateRepos = append(automateRepos, settingObj.Repos[r])
			}
		}
	}

	var wg sync.WaitGroup
	// Make a goroutine if commit is true
	if settingObj.SettingAddPullCommit.TimeMin > 0 {
		if settingObj.SettingAddPullCommit.TimeMin >= 1 {
			wg.Add(1)
			go CommandController(settingObj, dataUser[0].ForgitPath, automateRepos, "commit")
		}
	}

	//Make a goroutine if push is true
	if settingObj.SettingPush.TimeMin > 0 {
		if settingObj.SettingPush.TimeMin >= 1 {
			wg.Add(1)
			go CommandController(settingObj, dataUser[0].ForgitPath, automateRepos, "push")
		}
	}
	// This will make the program stay alive until go routines are done
	wg.Wait()
}
