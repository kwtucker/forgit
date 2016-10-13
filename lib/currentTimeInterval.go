package lib

import (
	"encoding/json"
	"io/ioutil"
	osuser "os/user"
)

// GetCurrentCPTimeMin gets the current time interval for either commit or push
func GetCurrentCPTimeMin(setObj Setting, gitCom string) (int, error) {
	var (
		fileu  []User
		cptime int
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
			switch gitCom {
			case "commit":
				cptime = fileu[0].Settings[s].SettingAddPullCommit.TimeMin
			case "push":
				cptime = fileu[0].Settings[s].SettingPush.TimeMin
			}
			break
		}
	}

	return cptime, err
}
