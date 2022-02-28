package grafana

import (
	"github.com/prometheus/prometheus/promql/parser"
)

// Dashboard is representing a json object for a grafana dashboard object
type Dashboard struct {
	Panels []Panel
}

// Panel is representing a json object for a grafana dashboard object
type Panel struct {
	Panels  []Panel
	Targets []Target
}

// Target is representing a json object for a grafana dashboard object
type Target struct {
	Expr string
}

// Unpack is recursivley gathering all the metrics from an array of panels
func Unpack(panels []Panel, metrics *[]string) {

	for _, panel := range panels {

		for _, target := range panel.Targets {

			expr, _ := parser.ParseExpr(target.Expr)

			selectors := parser.ExtractSelectors(expr)

			for _, selector := range selectors {

				for _, matcher := range selector {

					if matcher.Name == "__name__" {
						*metrics = append((*metrics), matcher.Value)
					}
				}
			}
		}
		Unpack(panel.Panels, metrics)
	}
}
