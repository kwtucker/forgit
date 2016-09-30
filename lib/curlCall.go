package lib

import (
	"io/ioutil"
	"net/http"
)

// Curforgit ...
func Curforgit(ghid string, fid string) ([]byte, error) {
	// Curl call that I am hooking up to forgit server later
	// resp, err := http.Get("http://whalebyte.com/api/users/" + ghid + "/" + fid)
	resp, err := http.Get("http://localhost:8100/api/users/" + ghid + "/" + fid)
	defer resp.Body.Close()
	databytes, err := ioutil.ReadAll(resp.Body)
	return databytes, err
}
