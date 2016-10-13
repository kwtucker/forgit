package lib

import (
	"io/ioutil"
	"net/http"
)

// Curlforgit ...
func Curlforgit(ghid string, fid string) ([]byte, error) {
	// Curl call that I am hooking up to forgit server later
	// resp, err := http.Get("http://whalebyte.com/api/users/" + ghid + "/" + fid)
	// resp, err := http.Get("http://localhost:8100/api/users/12258783/f892581c-edb5-4937-b10f-3624413059b1")
	resp, err := http.Get("http://localhost:8100/api/users/" + ghid + "/" + fid)
	defer resp.Body.Close()
	databytes, err := ioutil.ReadAll(resp.Body)
	return databytes, err
}
