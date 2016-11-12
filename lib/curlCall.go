package lib

import (
	"io/ioutil"
	"net/http"
)

// Curlforgit ...
func Curlforgit(init string, fid string) ([]byte, error) {
	// Calls forgitWeb and expects a full user object back.
	if init == "init" {
		resp, err := http.Get("http://forgit.whalebyte.com/api/users/" + fid + "/" + init)
		defer resp.Body.Close()
		databytes, err := ioutil.ReadAll(resp.Body)
		return databytes, err
	}
	// Only returns full user object if it has been updated.
	resp, err := http.Get("http://forgit.whalebyte.com/api/users/" + fid + "/no")
	defer resp.Body.Close()
	databytes, err := ioutil.ReadAll(resp.Body)
	return databytes, err
}
