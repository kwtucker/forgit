package lib

import (
	"io/ioutil"
	"net/http"
)

// Curlforgit ...
func Curlforgit(init string, fid string) ([]byte, error) {
	// Curl call that I am hooking up to forgit server later
	// resp, err := http.Get("http://forgit.whalebyte.com/api/users/" + ghid + "/" + fid)
	if init == "init" {
		resp, err := http.Get("http://localhost:8100/api/users/" + fid + "/" + init)
		defer resp.Body.Close()
		databytes, err := ioutil.ReadAll(resp.Body)
		return databytes, err
	}
	resp, err := http.Get("http://localhost:8100/api/users/" + fid + "/no")
	defer resp.Body.Close()
	databytes, err := ioutil.ReadAll(resp.Body)
	return databytes, err
}
