package exporter

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	muscleImbalancesLabels = []string{
		"agonist_muscle",
		"agonist_strength",
		"agonist_ratio",
		"antagonist_muscle",
		"antagonist_strength",
		"antagonist_ratio",
		"optimal_range_start",
		"optimal_range_end",
		"range_size",
		"body_region",
	}
	muscleImbalancesMetric = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "muscle_imbalance"),
		"Muscle imbalances",
		append(labels, muscleImbalancesLabels...),
		nil,
	)
)

func (e *EgymExporter) describeMuscleImbalanceMetrics(ch chan<- *prometheus.Desc) {
	ch <- muscleImbalancesMetric
}

func (e *EgymExporter) collectMuscleImbalanceMetrics(ch chan<- prometheus.Metric) {
	muscleImbalances, err := e.client.GetMuscleImbalances()
	if err != nil {
		log.Error("Could not retrieve muscle imbalances", err)
		return
	}

	for _, m := range *muscleImbalances {
		ch <- prometheus.MustNewConstMetric(
			muscleImbalancesMetric,
			prometheus.GaugeValue,
			m.Position,
			e.client.Username,
			m.AgonistMuscle,
			fmt.Sprintf("%f", m.AgonistStrengthValue),
			fmt.Sprintf("%f", m.AgonistRatio),
			m.AntagonistMuscle,
			fmt.Sprintf("%f", m.AntagonistStrengthValue),
			fmt.Sprintf("%f", m.AntagonistRatio),
			fmt.Sprintf("%f", m.OptimalRangeStartPosition),
			fmt.Sprintf("%f", m.OptimalRangeEndPosition),
			fmt.Sprintf("%f", m.RangeSize),
			m.BodyRegion,
		)
	}
}
