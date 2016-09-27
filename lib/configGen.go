package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	GithubID   int       `json:"githubID"`
	ForgitPath string    `json:"forgitPath"`
	Settings   []Setting `json:"settings,omitempty"`
}

// Setting ...
type Setting struct {
	SettingID int    `json:"setting_id"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	SettingNotifications
	SettingAddPullCommit
	SettingPush
	Repos []SettingRepo
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

// BuildConfig ...
func BuildConfig() {
	var (
		settingRepos        = []SettingRepo{}
		settings            = []Setting{}
		currentUserSettings = Setting{}
	)

	file, err := os.Open(*filename)
	if err != nil {
		if len(*filename) > 1 {
			fmt.Printf("Error: could not read config file %s.\n", *filename)
		}
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	// Overwrite the defaults
	if err := decoder.Decode(&c); err == io.EOF {
		fmt.Println(err)
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c.String())

	// Curl Call
	// resp, err := http.Get("https://api.github.com/users/kwtucker")
	// defer resp.Body.Close()
	// checkError(err)
	//
	// body, err := ioutil.ReadAll(resp.Body)
	// checkError(err)
	//
	// _, err = os.Stdout.Write(body)
	// checkError(err)

	for k := range settings {

	}
}
