package lib

import (
	"fmt"
	"io/ioutil"
	// "net/http"
	"encoding/json"
	"log"
	"os"
	osuser "os/user"
	"strconv"
	"time"
)

// Tells the user that the file exists and returns the config data
func fileExist(path string, forgitPath string, homeDir string) []byte {

	var (
		fileu []User
		// datau    []User
		dn       int64
		dateNow  string
		curldata []byte
		setdata  []Setting
		// setErr    []APIError
		// setUpdate []UpdateStatus
		update bool
	)

	// get unix time and convert it to a string for storage
	dn = time.Now().UTC().Unix()
	dateNow = strconv.FormatInt(dn, 10)

	//-=-=-=-=- TEMP to get test data from file -=-=-=-=-=-=-=
	// Read the config file in home dir
	// testfile, err := ioutil.ReadFile(homeDir + "/.forgitConfTest.json")
	// if err != nil {
	// 	os.Exit(1)
	// }
	// -=--=-=--END TEMP -=-=-=-=-=-=-=-

	existfile, err := ioutil.ReadFile(homeDir + "/.forgitConf.json")
	if err != nil {
		os.Exit(1)
	}
	// Set to user struct for local file
	json.Unmarshal(existfile, &fileu)

	// data from api
	// json.Unmarshal(testfile, &datau)

	// Call Forgit API for settings array []Setting
	gid := strconv.Itoa(fileu[0].GithubID)
	curldata, err = Curlforgit(gid, fileu[0].ForgitID)
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

	// parse the string timestamp to a int64 unix
	// fut, err := strconv.ParseInt(fileu[0].UpdateTime, 10, 64)
	// if err != nil {
	// 	fmt.Println("update time format error")
	// 	os.Exit(1)
	// }

	// dut, err := strconv.ParseInt(datau[0].UpdateTime, 10, 64)
	// if err != nil {
	// 	panic(err)
	// }

	// convert to a time.Time struct for comparing
	// fileUpdateTime := time.Unix(fut, 0)
	// dataUpdateTime := time.Unix(dut, 0)

	// if fileUpdateTime.After(dataUpdateTime) {
	if update {
		// Update the path in json
		fileu[0].ForgitPath = forgitPath
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
		return databytes
	}

	// if fileUpdateTime.Before(dataUpdateTime) {
	// 	// TODO: Need to sent post curl to api with new data and update mongodb on server
	//
	// 	//-=-=-=-=- TEMP to update test data from file -=-=-=-=-=-=-=
	// 	// Update the path in json
	// 	datau[0].ForgitPath = forgitPath
	// 	datau[0].UpdateTime = dateNow
	// 	// git byte array from MarshalIndent
	// 	databytes, err := json.MarshalIndent(datau, "", "    ")
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		os.Exit(1)
	// 	}
	// 	err = ioutil.WriteFile(path, databytes, 0644)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		os.Exit(1)
	// 	}
	// 	// -=--=-=--END TEMP -=-=-=-=-=-=-=-
	// 	return databytes
	// }
	return nil
}

// Creates a config file and puts server data to it.
func fileNotExist(path string, j []byte, forgitPath string) {

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
	if _, err = os.Stat(forgitPath + "Forgit/"); os.IsNotExist(err) {
		err = os.Mkdir(forgitPath+"Forgit/", 0700)
		if err != nil {
			fmt.Println(err)
		}
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
		now := time.Now().UTC().Unix()
		nowString := strconv.FormatInt(now, 10)
		u[0].UpdateTime = nowString

		// git byte array from MarshalIndent
		databytes, err = json.MarshalIndent(u, "", "    ")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fileNotExist(homeDir.HomeDir+"/.forgitConf.json", databytes, forgitPath)
	}

	// File Exists Print
	fileExist(homeDir.HomeDir+"/.forgitConf.json", forgitPath, homeDir.HomeDir)

	fmt.Println("\n\tYour Config is in --> " + homeDir.HomeDir + "/.forgitConf.json\n")

}
