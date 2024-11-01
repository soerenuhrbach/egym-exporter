package exporter

import (
	"slices"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const bodyNamespace = "body"

var (
	bodyLabels = []string{"type", "source", "source_label", "unit"}

	bodyMetric = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, bodyNamespace, "metrics"),
		"Measured body metric",
		append(labels, bodyLabels...),
		nil,
	)

	validUnits = []string{"PERCENT", "PERCENTS", "KG", "LITER", "ANGLE", "KJ", "CM"}
)

func (e *EgymExporter) describeBodyMetrics(ch chan<- *prometheus.Desc) {
	ch <- bodyMetric
}

func (e *EgymExporter) collectBodyMetrics(ch chan<- prometheus.Metric) {
	metrics, err := e.client.GetBodyMetrics()
	if err != nil {
		log.Error("Could not fetch body metrics", err)
		return
	}

	for _, m := range *metrics {
		if strings.Contains(m.Type, "LOW") || strings.Contains(m.Type, "TOP") {
			continue
		}

		var unit string
		metricNamePartials := strings.Split(m.Type, "_")

		for i, p := range metricNamePartials {
			if slices.Contains(validUnits, p) {
				metricNamePartials = slices.Delete(metricNamePartials, i, i+1)
				unit = p
				break
			}
		}

		metricNameWithoutUnit := strings.Join(metricNamePartials, "_")

		ch <- prometheus.MustNewConstMetric(
			bodyMetric,
			prometheus.CounterValue,
			m.Value,
			e.client.Username,
			metricNameWithoutUnit,
			m.Source,
			m.SourceLabel,
			unit,
		)
	}
}
