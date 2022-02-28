package util

import (
	"io/ioutil"
	"net/http"
)

// ReadFile is given a path returns bytes
func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	CheckError(err)
	return data
}

// DownloadFile is given a url returns bytes
func DownloadFile(url string) []byte {
	response, err := http.Get(url)
	CheckError(err)
	data, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	return data
}

// CheckError is checking for non nil error and panics
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Distinct is taking an existing array of strings and returns a distinct array of strings
func Distinct(values []string) []string {

	keys := make(map[string]bool)

	uniqueList := make([]string, 0, 20)

	for _, value := range values {
		if _, exists := keys[value]; !exists {
			keys[value] = true
			uniqueList = append(uniqueList, value)
		}
	}
	return uniqueList
}
