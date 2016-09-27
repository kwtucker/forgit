package lib

import (
	"fmt"
	"io/ioutil"
	// "net/http"
	"encoding/json"
	// "io/ioutil"
	"os"
	osuser "os/user"
)

// User ..
type User struct {
	GithubID   int       `json:"githubID"`
	ForgitPath string    `json:"forgitPath"`
	UpdateTime string    `json:"updateTime"`
	Settings   []Setting `json:"settings,omitempty"`
}

// Setting ...
type Setting struct {
	SettingID            int            `json:"setting_id"`
	Name                 string         `json:"name"`
	Status               int            `json:"status"`
	SettingNotifications map[string]int `json:"notifications"`
	SettingAddPullCommit map[string]int `json:"addPullCommit"`
	SettingPush          map[string]int `json:"push"`
	Repos                []SettingRepo  `json:"repos"`
}

// SettingNotifications ...
type SettingNotifications struct {
	Status   int `json:"status"`
	OnError  int `json:"onError"`
	OnCommit int `json:"onCommit"`
	OnPush   int `json:"onPush"`
}

// SettingAddPullCommit ...
type SettingAddPullCommit struct {
	Status  int `json:"status"`
	TimeMin int `json:"timeMinute"`
}

// SettingPush ...
type SettingPush struct {
	Status  int `json:"status"`
	TimeMin int `json:"timeMinute"`
}

// SettingRepo ...
type SettingRepo struct {
	GithubRepoID *int    `json:"github_repo_id"`
	Name         *string `json:"name"`
	Status       int     `json:"status"`
}

// Tells the user that the file exists and returns the config data
func fileExist(path string) []byte {
	var (
		u []User
	)

	// Read the config file in home dir
	file, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(1)
	}

	// Set to user struct
	json.Unmarshal(file, &u)
	filebytes, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return filebytes
}

// Creates a config file and puts server data to it.
func fileNotExist(path string, j []byte) {
	var (
		f     *os.File
		err   error
		jsonU []User
	)

	// set data to the user struct and indent format
	json.Unmarshal(j, &jsonU)
	filebytes, err := json.MarshalIndent(jsonU, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create the file and defer close
	f, err = os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	// defer file close tell end of function
	defer f.Close()

	// Write to file and not set a var
	_, err = f.Write(filebytes)
	if err != nil {
		fmt.Println(err)
	}
	// save
	f.Sync()
	fmt.Println("created")
}

// BuildConfig ...
func BuildConfig(forgitPath string) {
	// Get Home Directory
	homeDir, err := osuser.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//IF config file doesn't exist. Create it
	if _, err := os.Stat(homeDir.HomeDir + "/.forgitConf2.json"); os.IsNotExist(err) {
		// // Curl call that I am hooking up to forgit server later
		// resp, err := http.Get("https://api.github.com/users/kwtucker")
		// defer resp.Body.Close()
		// checkError(err)
		//
		// databytes, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		//   fmt.Println(err)
		// 	os.Exit(1)
		// })
		// fileNotExist(homeDir.HomeDir+"/.forgitConf2.json", databytes)

		// NOTE: need to replace p with curl call
		p := fileExist(homeDir.HomeDir + "/.forgitConf.json")
		fileNotExist(homeDir.HomeDir+"/.forgitConf2.json", p)

	}

	// File Exists Print
	p := fileExist(homeDir.HomeDir + "/.forgitConf.json")
	fmt.Println("Your Config already exists in --> " + homeDir.HomeDir)
	fmt.Println(string(p))

}
