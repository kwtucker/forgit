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

// ConfigFileReadUpdate Tells the user that the file exists and returns the config data
func ConfigFileReadUpdate(path string, forgitPath string, homeDir string, uuid string, reqt string) {
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

	// Check if the json file is there.
	existfile, err := ioutil.ReadFile(homeDir + "/.forgitConf.json")
	if err != nil {
		fmt.Println(".forgitConf.json Does not exist, Re-download Forgit")
		os.Exit(1)
	}

	// Set to user struct for local file
	json.Unmarshal(existfile, &fileu)

	// Check internet and if so get user data from forgit web.
	isInternet := InternetCheck()
	if isInternet {
		curldata, err = Curlforgit(reqt, uuid)
		if err != nil {
			log.Println(err)
		}
		// The response for a bad request will always be 42 bytes
		if len(curldata) == 42 {
			fmt.Println("Bad UUID credentials,")
			fmt.Println(" 1. Try forgit init again and make sure to copy all the UUID from the dashboard on your browser. http://forgit.whalebyte.com/dashboard/")
			fmt.Println(" 2. If you did not get the CLI you are using from forgit.whalebyte.com, be sure to \nlogin to forgit.whalebyte.com and get your own UUID from the dashboard.")
			os.Exit(1)
		} else {
			fmt.Println("Curl data on init returned something unexpected.\nTry forgit init again.")
		}

		// If the data returned is large update the settings.
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
}

// Creates a config file and puts server data to it.
func createConfigFileIfNotExist(homeDir string, uuid string, forgitPath string) {
	var (
		f    *os.File
		err  error
		uarr []User
	)

	// Create a temp user.
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

	// Put user in a array and unpack the user.
	uarr = append(uarr, u)
	filebytes, err := json.MarshalIndent(uarr, "", "    ")
	if err != nil {
		log.Println(err)
	}

	// If the Forgit directory doesn't exist make it.
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

	ConfigFileReadUpdate(homeDir+"/.forgitConf.json", forgitPath, homeDir, uuid, "init")
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
		createConfigFileIfNotExist(homeDir.HomeDir, uuid, forgitPath)
	}

	ConfigFileReadUpdate(homeDir.HomeDir+"/.forgitConf.json", forgitPath, homeDir.HomeDir, uuid, "init")

	fmt.Println("\n\tYour Config Is In --> " + homeDir.HomeDir + "/.forgitConf.json\n")
	fmt.Println("\n\tNow Run --> " + "forgit start\n")

}
