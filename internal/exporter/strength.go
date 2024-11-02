package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	strengthLabels = []string{"source", "source_label", "body_region", "muscle", "exercise", "exercise_code", "progress"}

	strengthMetrics = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "strength"),
		"Strength metrics",
		append(labels, strengthLabels...),
		nil,
	)
)

func (e *EgymExporter) describeStrengthMetrics(ch chan<- *prometheus.Desc) {
	ch <- strengthMetrics
}

func (e *EgymExporter) collectStrengthMetrics(ch chan<- prometheus.Metric) {
	strengths, err := e.client.GetStrengthMetrics()
	if err != nil {
		log.Error("could not retrieve strength metrics", err)
		return
	}

	for _, s := range *strengths {
		ch <- prometheus.MustNewConstMetric(
			strengthMetrics,
			prometheus.GaugeValue,
			s.Strength.Value,
			e.client.Username,
			s.Source,
			s.SourceLabel,
			s.BodyRegion,
			e.client.GetMuscleFromStrengthExercise(s.Exercise.Code),
			s.Exercise.Code,
			s.Exercise.Label,
			s.Strength.Progress,
		)
	}
}
