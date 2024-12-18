package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/soerenuhrbach/egym-exporter/internal/egym"
)

const namespace = "egym"

var (
	labels = []string{"user"}
)

type EgymExporter struct {
	client *egym.EgymClient
}

func (e *EgymExporter) Describe(ch chan<- *prometheus.Desc) {
	e.describeBioAgeMetrics(ch)
	e.describeActivityLevelMetrics(ch)
	e.describeBodyMetrics(ch)
	e.describeStrengthMetrics(ch)
	e.describeFlexibilityMetrics(ch)
	e.describeMuscleImbalanceMetrics(ch)
	e.describeExerciseMetrics(ch)
}

func (e *EgymExporter) Collect(ch chan<- prometheus.Metric) {
	e.collectBioAgeMetrics(ch)
	e.collectActivityLevelMetrics(ch)
	e.collectBodyMetrics(ch)
	e.collectStrengthMetrics(ch)
	e.collectFlexibilityMetrics(ch)
	e.collectMuscleImbalanceMetrics(ch)
	e.collectExerciseMetrics(ch)
}

func NewEgymExporter(client *egym.EgymClient) *EgymExporter {
	return &EgymExporter{
		client: client,
	}
}
