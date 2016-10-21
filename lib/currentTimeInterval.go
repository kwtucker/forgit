package lib

import (
	"encoding/json"
	"io/ioutil"
	osuser "os/user"
)

// GetCurrentCPTimeMin gets the current time interval for either commit or push
func GetCurrentCPTimeMin(setObj Setting, gitCom string) (int, int, int, int, error) {
	var (
		fileu                         []User
		cptime                        int
		noteerr, notepush, notecommit int
	)
	// Home directory of the user machine
	homeDir, err := osuser.Current()
	// config file
	configFile, err := ioutil.ReadFile(homeDir.HomeDir + "/.forgitConf.json")
	// Set to user struct for local file
	json.Unmarshal(configFile, &fileu)

	// Grab the settings object from config file
	// set the time interval to the commit/push var cptime
	for s := range fileu[0].Settings {
		if fileu[0].Settings[s].Name == setObj.Name {
			noteerr = fileu[0].Settings[s].OnError
			notepush = fileu[0].Settings[s].OnPush
			notecommit = fileu[0].Settings[s].OnCommit
			switch gitCom {
			case "commit":
				// cptime = fileu[0].Settings[s].SettingAddPullCommit.TimeMin
				cptime = fileu[0].Settings[s].SettingAddPullCommit.TimeMin

			case "push":
				// cptime = fileu[0].Settings[s].SettingPush.TimeMin
				cptime = fileu[0].Settings[s].SettingPush.TimeMin

			}
			break
		}
	}

	if setObj.Name == "fgtDefault" {
		noteerr = 1
		notecommit = 1
		notepush = 1
	}

	return cptime, noteerr, notecommit, notepush, err
}
