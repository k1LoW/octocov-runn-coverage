package runncoverage

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k1LoW/octocov/report"
	"github.com/k1LoW/runn"
)

func TestConverter(t *testing.T) {
	tests := []struct {
		in   *runn.Coverage
		want *report.CustomMetricSet
	}{
		{
			&runn.Coverage{
				Specs: []*runn.SpecCoverage{
					{
						Key: "testapi:0.0.1",
						Coverages: map[string]int{
							"GET /api/v1/users":      1,
							"GET /api/v1/users/{id}": 0,
							"POST /api/v1/users":     1,
							"PUT /api/v1/users/{id}": 0,
						},
					},
				},
			},
			&report.CustomMetricSet{
				Key:  "runn coverage",
				Name: "Runbook Coverage Report",
				Metrics: []*report.CustomMetric{
					{
						Key:   "testapi:0.0.1",
						Name:  "testapi:0.0.1",
						Value: 50.0,
						Unit:  "%",
					},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, err := Convert(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			opts := []cmp.Option{
				cmpopts.IgnoreFields(report.CustomMetricSet{}, "report"),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Error(diff)
			}
		})
	}
}
