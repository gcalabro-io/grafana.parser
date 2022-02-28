package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gmcalab/parser/grafana"
	"github.com/gmcalab/parser/util"
)

func main() {

	if len(os.Args) <= 1 {
		usage()
	}

	var dashboard grafana.Dashboard
	var data []byte

	source := os.Args[1]

	if source == "-h" {
		help()
		os.Exit(0)
	}

	if source != "-p" && source != "-u" {
		usage()
	}

	target := os.Args[2]

	if source == "-p" {
		data = util.ReadFile(target)
	} else {
		response, err := http.Get(target)
		util.CheckError(err)
		data, err = ioutil.ReadAll(response.Body)
		util.CheckError(err)
	}

	json.Unmarshal(data, &dashboard)

	metrics := make([]string, 0, 20)

	grafana.Unpack(dashboard.Panels, &metrics)

	uniqueMetrics := util.Distinct(metrics)

	fmt.Printf("\n\n------------------------\n")
	fmt.Printf("Total Metrics Found: %d\n", len(metrics))
	fmt.Printf("Total Unique Metrics: %d\n", len(uniqueMetrics))
	fmt.Printf("------------------------\n\n")
	for _, metric := range uniqueMetrics {
		fmt.Println(metric)
	}
}

func help() {
	fmt.Println("-h\t help")
	fmt.Println("-p <path>\t file path for dashboard json")
	fmt.Println("-u <url>\t url for dashboard json")
}

func usage() {
	panic("Incorrect usage, please use -h for help")
}
