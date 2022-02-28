package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gcalabro-io/parser/grafana"
	"github.com/gcalabro-io/parser/util"
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
		data = util.DownloadFile(target)
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

// help is showing the help menu
func help() {
	fmt.Println("-h\t help")
	fmt.Println("-p <path>\t file path for dashboard json")
	fmt.Println("-u <url>\t url for dashboard json")
}

// usage is showing the panic message for incorrect usage
func usage() {
	panic("Incorrect usage, please use -h for help")
}
