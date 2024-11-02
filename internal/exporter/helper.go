package exporter

import (
	"slices"
	"strings"
)

var validUnits = []string{"PERCENT", "PERCENTS", "KG", "LITER", "ANGLE", "KJ", "CM", "DEGREE"}

func parseUnitAndNameFromMetricType(metricType string) (name string, unit string) {
	metricNamePartials := strings.Split(metricType, "_")

	for i, p := range metricNamePartials {
		if slices.Contains(validUnits, p) {
			metricNamePartials = slices.Delete(metricNamePartials, i, i+1)
			unit = p
			break
		}
	}

	name = strings.Join(metricNamePartials, "_")
	return
}
