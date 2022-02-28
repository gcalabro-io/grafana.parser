package util

import (
	"io/ioutil"
)

func ReadFile(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	CheckError(err)
	return data
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Distinct(values []string) []string {

	keys := make(map[string]bool)

	uniqueList := make([]string, 0, 20)

	for _, value := range values {
		if _, v := keys[value]; !v {
			keys[value] = true
			uniqueList = append(uniqueList, value)
		}
	}
	return uniqueList
}
