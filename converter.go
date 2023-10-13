package runncoverage

import (
	"github.com/k1LoW/octocov/report"
	"github.com/k1LoW/runn"
)

func Convert(cov *runn.Coverage) (*report.CustomMetricSet, error) {
	cset := &report.CustomMetricSet{
		Key:  "runn coverage",
		Name: "Runbook Coverage Report",
	}
	for _, spec := range cov.Specs {
		var total, covered int
		for _, c := range spec.Coverages {
			total++
			if c > 0 {
				covered++
			}
		}
		cset.Metrics = append(cset.Metrics, &report.CustomMetric{
			Key:   spec.Key,
			Name:  spec.Key,
			Value: float64(covered) / float64(total) * 100,
			Unit:  "%",
		})
	}
	return cset, nil
}
