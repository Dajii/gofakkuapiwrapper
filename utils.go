package gofakkuapiwrapper

import (
	"io/ioutil"
	"net/http"
	"net/url"
)


func loadJSON(url *url.URL) (string, error) {
	resp, getErr := http.Get(url.String())
	if getErr != nil {
		return "", getErr
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return string(body), readErr
	}
	return string(body), nil
}
