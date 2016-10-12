package lib

// APIError ...
type APIError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//UpdateStatus ...
type UpdateStatus struct {
	Update string `json:"update"`
}

// User ..
type User struct {
	GithubID   int       `json:"githubID"`
	ForgitID   string    `json:"forgitID"`
	ForgitPath string    `json:"forgitPath"`
	UpdateTime string    `json:"updateTime"`
	Settings   []Setting `json:"settings,omitempty"`
}

// Setting ...
type Setting struct {
	Name                 string `json:"name"`
	Status               int    `json:"status"`
	SettingNotifications `json:"notifications"`
	SettingAddPullCommit `json:"addPullCommit"`
	SettingPush          `json:"push"`
	Repos                []SettingRepo `json:"repos"`
}

// SettingNotifications ...
type SettingNotifications struct {
	OnError  int `json:"onError"`
	OnCommit int `json:"onCommit"`
	OnPush   int `json:"onPush"`
}

// SettingAddPullCommit ...
type SettingAddPullCommit struct {
	TimeMin int `json:"timeMinute"`
}

// SettingPush ...
type SettingPush struct {
	TimeMin int `json:"timeMinute"`
}

// SettingRepo ...
type SettingRepo struct {
	GithubRepoID int    `json:"github_repo_id"`
	Name         string `json:"name"`
	Status       int    `json:"status"`
}
