package grafana

import (
	"github.com/prometheus/prometheus/promql/parser"
)

type Dashboard struct {
	Panels []Panel
}

type Panel struct {
	Panels  []Panel
	Targets []Target
}

type Target struct {
	Expr string
}

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
