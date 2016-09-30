package lib

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
