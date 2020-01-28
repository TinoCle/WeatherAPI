// package myip is designed to return YOUR IP Address
// and that's it.
package myip

import (
	"io/ioutil"
	"net/http"
)

var sites = []string{
	"http://icanhazip.com",
	"http://canihazip.com/s",
}

// GetMyIP iterates through a list of websites and returns
// when it gets a successful response.
func GetMyIP() (string, error) {
	var err error

	// Keep looping until we actually get a real answer
	for _, site := range sites {
		resp, e := http.Get(site)
		if e != nil {
			err = e
			continue
		}

		defer resp.Body.Close()
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			err = e
			continue
		}

		return string(body), nil
	}

	return "", err
}
