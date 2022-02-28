package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gmcalab/parser/grafana"
	"github.com/gmcalab/parser/util"
)

func main() {

	dashboardYaml := os.Args[1]

	metrics := make([]string, 0, 20)

	var dashboard grafana.Dashboard

	data := util.ReadFile(dashboardYaml)

	json.Unmarshal(data, &dashboard)

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
