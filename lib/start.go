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
	"strconv"
	"sync"
	"time"
)

// settingGroupsCheck checks if the user object has the matching group name.
func settingGroupsCheck(groupName string, dataUser User) (Setting, bool) {
	for u := range dataUser.Settings {
		if dataUser.Settings[u].Name == groupName {
			return dataUser.Settings[u], true
		}
	}
	return dataUser.Settings[0], false
}

// ForgitDirReposNames reads the Forgit directory, then checks if those
// directories are git repos by looking for the .git dir.
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

// InternetCheck Checks if internet exists by calling google.
func InternetCheck() bool {
	_, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("\nNo internet")
		return false
	}
	return true
}

// Start is the main controller for the app that:
// - Calls the forgit server.
// - Checks the internet.
// - Dispatches commands based on the request from user.
func Start(c *cli.Context) {

	var (
		commit, push int
		group        string
		dataUser     []User
		settingObj   Setting
		settingRepos []SettingRepo
		settingRepo  SettingRepo
		setExist     bool
		dn           int64
		dateNow      string
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
		fmt.Println("* Please first run --> forgit init ")
		fmt.Println()
		return
	}

	// Read .forgitConf.json file
	// data from configfile
	configfile, err := ioutil.ReadFile(homeDir.HomeDir + "/.forgitConf.json")
	if err != nil {
		os.Exit(1)
	}
	json.Unmarshal(configfile, &dataUser)

	// If the internetConnection is true
	if internetConnection {
		var curldata []byte
		// Grab most recent data and set it to the datauser
		curldata, err = Curlforgit("init", dataUser[0].ForgitID)
		if err != nil {
			log.Println(err)
		}
		var setdata []Setting
		json.Unmarshal(curldata, &setdata)
		dataUser[0].Settings = setdata
	}

	// update time and format it to unix
	dn = time.Now().UTC().Unix()
	dateNow = strconv.FormatInt(dn, 10)
	dataUser[0].UpdateTime = dateNow

	// Unpack the json data.
	databytes, err := json.MarshalIndent(dataUser, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Update the config file.
	err = ioutil.WriteFile(homeDir.HomeDir+"/.forgitConf.json", databytes, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("\nThis session will have the following settings:")

	// If group exists grab group name and set the values
	if c.Args().First() != "" {
		group = c.Args().First()
		settingObj, setExist = settingGroupsCheck(group, dataUser[0])
		if setExist == false {
			fmt.Println()
			fmt.Println("* Did Not Start!")
			fmt.Println("* Setting group does not exist.")
			fmt.Println()
			return
		}
	} else {
		group = "-1"
	}

	// If commit set value.
	if c.IsSet("commit") && c.Int("commit") != 0 && group == "-1" {
		commit = c.Int("commit")
	} else {
		commit = -1
	}

	// If push set value.
	if c.IsSet("push") && c.Int("push") != 0 && group == "-1" {
		push = c.Int("push")
	} else {
		push = -1
	}

	// Get byte array from MarshalIndent
	_, err = json.MarshalIndent(dataUser, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// If no push,commit, or group set. Grab setting with status of 1
	if c.IsSet("push") == false && c.IsSet("commit") == false && group == "-1" {
		for s := range dataUser[0].Settings {
			if dataUser[0].Settings[s].Status == 1 {
				settingObj = dataUser[0].Settings[s]
			}
		}
	}

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
		repoNames := ForgitDirReposNames(dataUser[0].ForgitPath)
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
			Name:   "forgitDefault",
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
		automateRepos = settingRepos
	}

	fmt.Println("\nSetting Name: ", settingObj.Name)
	fmt.Println("Commit Time: ", settingObj.SettingAddPullCommit.TimeMin)
	fmt.Println("Push Time: ", settingObj.SettingPush.TimeMin)
	fmt.Println()

	// If the length of repos in the Forgit dir is 0 stop the app.
	if len(automateRepos) == 0 {
		log.Println(": You don't have any repos to automate.\n" +
			"\tOr you don't have any selected in setting group.\n" +
			"\tSelect repos in the " + settingObj.Name + " setting group and restart. forgit start")
		os.Exit(1)
	}

	// Create a WaitGroup for the go routines
	var wg sync.WaitGroup

	// Make a goroutine if commit is true
	if settingObj.SettingAddPullCommit.TimeMin > 0 {
		if settingObj.SettingAddPullCommit.TimeMin >= 1 {
			wg.Add(1)
			go CommandController(settingObj, dataUser[0].ForgitPath, automateRepos, dataUser[0].ForgitID, internetConnection, "commit")
		}
	}

	//Make a goroutine if push is true
	if settingObj.SettingPush.TimeMin > 0 {
		if settingObj.SettingPush.TimeMin >= 1 && internetConnection == true {
			wg.Add(1)
			go CommandController(settingObj, dataUser[0].ForgitPath, automateRepos, dataUser[0].ForgitID, internetConnection, "push")
		}
	}
	// This will make the program stay alive until go routines are done
	wg.Wait()
}
