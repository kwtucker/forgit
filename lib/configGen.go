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
			fileu[0].ForgitPath = forgitPath + "Forgit/"
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
func fileNotExist(homeDir string, j []byte, forgitPath string) {

	var (
		f   *os.File
		err error
	)

	u := User{
		ForgitID:   "",
		ForgitPath: "",
		UpdateTime: "",
		Settings:   []Setting{},
	}
	filebytes, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		log.Println(err)
	}

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

		var (
			u         []User
			file      []byte
			databytes []byte
		)

		// Read the config file in home dir
		file, err = ioutil.ReadFile(homeDir.HomeDir + "/.forgitConfTest.json")
		if err != nil {
			os.Exit(1)
		}

		// Set to user struct for local file
		json.Unmarshal(file, &u)
		// Update the path in json
		u[0].ForgitPath = forgitPath
		u[0].ForgitID = uuid
		now := time.Now().UTC().Unix()
		nowString := strconv.FormatInt(now, 10)
		u[0].UpdateTime = nowString

		// git byte array from MarshalIndent
		databytes, err = json.MarshalIndent(u, "", "    ")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fileNotExist(homeDir.HomeDir, databytes, forgitPath)
	}

	// File Exists Print
	FileExist(homeDir.HomeDir+"/.forgitConf.json", forgitPath, homeDir.HomeDir, uuid, "init")

	fmt.Println("\n\tYour Config is in --> " + homeDir.HomeDir + "/.forgitConf.json\n")

}
