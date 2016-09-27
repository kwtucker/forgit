package lib

import (
	"fmt"
	"io/ioutil"
	// "net/http"
	"encoding/json"
	"github.com/google/go-github/github"
	"io/ioutil"
	"os"
	osuser "os/user"
	"time"
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

func fileExist(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(1)
		// return
	}
	var u []User
	json.Unmarshal(file, &u)
	filebytes, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(filebytes)
}

func notExist(path string) {
	var (
		f            *os.File
		err          error
		settingRepos = []SettingRepo{}
		settings     = []Setting{}
	)

	f, err = os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// set the time
	for s := range settings {
		// Create json
		for r := range settings[s].Repos {
			userSettingsRepo := SettingRepo{
				GithubRepoID: settings[s].Repos[s].ID,
				Name:         settings[s].Repos[s].Name,
				Status:       settings[s].Repos[s].Status,
			}
			settingRepos = append(settingRepos, userSettingsRepo)
		}

		userSetting := Setting{
			SettingID: settings[s].ID,
			Name:      settings[s].Name,
			Status:    settings[s].Status,
			SettingNotifications: SettingNotifications{
				Status:   settings[s].SettingNotifications.Status,
				OnError:  settings[s].SettingNotifications.OnError,
				OnCommit: settings[s].SettingNotifications.OnCommit,
				OnPush:   settings[s].SettingNotifications.OnPush,
			},
			SettingAddPullCommit: SettingAddPullCommit{
				Status:  settings[s].SettingAddPullCommit.Status,
				TimeMin: settings[s].SettingAddPullCommit.TimeMin,
			},
			SettingPush: SettingPush{
				Status:  settings[s].SettingPush.Status,
				TimeMin: settings[s].SettingPush.TimeMin,
			},
			Repos: settingRepos,
		}
		settings = append(settings, userSetting)
	}

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	timenow := &github.Timestamp{time.Now().In(location)}
	user := &User{
		GithubID:   12345,
		ForgitPath: fPath,
		UpdateTime: timenow.String(),
		Settings:   settings,
	}
	// move file

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err = f.Write(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("wrote %d bytes\n", n2)
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
	if _, err := os.Stat(homeDir.HomeDir + "/.forgitConf.json"); os.IsNotExist(err) {
		fileNotExist(homeDir.HomeDir + "/.forgitConf.json")
		fmt.Println("Nope")
		os.Exit(1)
	}

	// File Exists
	p := fileExist(homeDir.HomeDir + "/.forgitConf.json")
	fmt.Println(p)

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
}
