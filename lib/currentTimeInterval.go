package lib

import (
	"encoding/json"
	"io/ioutil"
	osuser "os/user"
)

func GetCurrentCPTimeMin(setObj Setting, gitCom string) (int, error) {
	var (
		fileu  []User
		cptime int
	)
	homeDir, err := osuser.Current()
	existfile, err := ioutil.ReadFile(homeDir.HomeDir + "/.forgitConf.json")
	// Set to user struct for local file
	json.Unmarshal(existfile, &fileu)
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
