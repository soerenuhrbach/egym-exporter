package exporter

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	bodyLabels = []string{"type", "source", "source_label", "unit"}

	bodyMetric = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "body"),
		"Measured body metric",
		append(labels, bodyLabels...),
		nil,
	)
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

		name, unit := parseUnitAndNameFromMetricType(m.Type)

		ch <- prometheus.MustNewConstMetric(
			bodyMetric,
			prometheus.CounterValue,
			m.Value,
			e.client.Username,
			name,
			m.Source,
			m.SourceLabel,
			unit,
		)
	}
}
