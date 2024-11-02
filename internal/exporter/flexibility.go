package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	flexibilityLabels = []string{"type", "source", "source_label", "interpretation", "unit"}

	flexibilityMetric = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "flexibility"),
		"Flexibility metrics",
		append(labels, flexibilityLabels...),
		nil,
	)
)

func (e *EgymExporter) describeFlexibilityMetrics(ch chan<- *prometheus.Desc) {
	ch <- flexibilityMetric
}

func (e *EgymExporter) collectFlexibilityMetrics(ch chan<- prometheus.Metric) {
	metrics, err := e.client.GetFlexibilityMetrics()
	if err != nil {
		log.Error("Could not retrieve flexibility metrics")
		return
	}

	for _, m := range *metrics {

		name, unit := parseUnitAndNameFromMetricType(m.Type)

		ch <- prometheus.MustNewConstMetric(
			flexibilityMetric,
			prometheus.GaugeValue,
			m.Value,
			e.client.Username,
			name,
			m.Source,
			m.SourceLabel,
			m.ValueInterpretation,
			unit,
		)
	}
}
