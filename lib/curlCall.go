package lib

import (
	"io/ioutil"
	"net/http"
)

// Curlforgit ...
func Curlforgit(init string, fid string) ([]byte, error) {
	// Curl call that I am hooking up to forgit server later
	if init == "init" {
		resp, err := http.Get("http://forgit-stage.whalebyte.com/api/users/" + fid + "/" + init)
		defer resp.Body.Close()
		databytes, err := ioutil.ReadAll(resp.Body)
		return databytes, err
	}
	resp, err := http.Get("http://forgit-stage.whalebyte.com/api/users/" + fid + "/no")
	defer resp.Body.Close()
	databytes, err := ioutil.ReadAll(resp.Body)
	return databytes, err
}
