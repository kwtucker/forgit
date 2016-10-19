package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	osuser "os/user"
	"strconv"
	"time"
)

// FileExist Tells the user that the file exists and returns the config data
// func FileExist(path string, forgitPath string, homeDir string) []byte {
func FileExist(path string, forgitPath string, homeDir string, uuid string, reqt string) {
	var (
		fileu    []User
		dn       int64
		dateNow  string
		curldata []byte
		setdata  []Setting
		update   bool
	)
	// get unix time and convert it to a string for storage
	dn = time.Now().UTC().Unix()
	dateNow = strconv.FormatInt(dn, 10)

	existfile, err := ioutil.ReadFile(homeDir + "/.forgitConf.json")
	if err != nil {
		fmt.Println(".forgitConf.json Does not exist, Re-download Forgit")
		os.Exit(1)
	}

	// Set to user struct for local file
	json.Unmarshal(existfile, &fileu)

	curldata, err = Curlforgit(reqt, uuid)
	if err != nil {
		log.Println(err)
	}

	if len(curldata) > 200 {
		// Format curl data and set it to settings array
		err = json.Unmarshal(curldata, &setdata)
		fileu[0].Settings = setdata
		update = true
	} else {
		update = false
	}

	if update {
		// Update the path in json
		if reqt == "init" {
			if _, err = os.Stat(forgitPath + "Forgit/"); os.IsNotExist(err) {
				err = os.Mkdir(forgitPath+"Forgit/", 0700)
				if err != nil {
					fmt.Println(err)
				}
			}
			fileu[0].ForgitPath = forgitPath + "Forgit/"
			fileu[0].ForgitID = uuid
		} else {
			fileu[0].ForgitPath = forgitPath
		}

		fileu[0].UpdateTime = dateNow

		// git byte array from MarshalIndent
		databytes, err := json.MarshalIndent(fileu, "", "    ")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Write to file with updated info
		err = ioutil.WriteFile(path, databytes, 0644)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

// Creates a config file and puts server data to it.
// func fileNotExist(homeDir string, j []byte, forgitPath string) {
func fileNotExist(homeDir string, uuid string, forgitPath string) {
	var (
		f    *os.File
		err  error
		uarr []User
	)

	u := User{
		ForgitID:   uuid,
		ForgitPath: forgitPath,
		UpdateTime: "0",
		Settings: []Setting{
			Setting{
				Name:   "General",
				Status: 1,
				SettingNotifications: SettingNotifications{
					OnError:  1,
					OnCommit: 1,
					OnPush:   1,
				},
				SettingAddPullCommit: SettingAddPullCommit{
					TimeMin: 2,
				},
				SettingPush: SettingPush{
					TimeMin: 60,
				},
				Repos: []SettingRepo{
					SettingRepo{
						GithubRepoID: 0,
						Name:         "r",
						Status:       0,
					},
				},
			},
		},
	}
	uarr = append(uarr, u)
	filebytes, err := json.MarshalIndent(uarr, "", "    ")
	if err != nil {
		log.Println(err)
	}
	// f, err = os.Create(homeDir + ".forgitConf.json")
	// f.Close()
	if _, err = os.Stat(forgitPath + "Forgit/"); os.IsNotExist(err) {
		err = os.Mkdir(forgitPath+"Forgit/", 0700)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Create the file and defer close
	f, err = os.Create(homeDir + "/.forgitConf.json")
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

	FileExist(homeDir+"/.forgitConf.json", forgitPath, homeDir, uuid, "init")
}

// BuildConfig ...
func BuildConfig(forgitPath string, uuid string) {
	// Get Home Directory
	homeDir, err := osuser.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// If config file doesn't exist. Create it
	if _, err = os.Stat(homeDir.HomeDir + "/.forgitConf.json"); os.IsNotExist(err) {
		fileNotExist(homeDir.HomeDir, uuid, forgitPath)
	}

	// File Exists Print
	FileExist(homeDir.HomeDir+"/.forgitConf.json", forgitPath, homeDir.HomeDir, uuid, "init")

	fmt.Println("\n\tYour Config Is In --> " + homeDir.HomeDir + "/.forgitConf.json\n")
	fmt.Println("\n\tNow Run --> " + "fgt start\n")

}
