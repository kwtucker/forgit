package lib

// APIError ...
type APIError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Message ...
type Message struct {
	Title string
	Body  string
}

//UpdateStatus ...
type UpdateStatus struct {
	Update string `json:"update"`
}

// User ..
type User struct {
	// GithubID   int       `json:"githubID"`
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
	TimeMin int `json:"timeMin"`
}

// SettingPush ...
type SettingPush struct {
	TimeMin int `json:"timeMin"`
}

// SettingRepo ...
type SettingRepo struct {
	GithubRepoID int    `json:"githubrepoid"`
	Name         string `json:"name"`
	Status       int    `json:"status"`
}
