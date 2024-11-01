package exporter

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const activityNamespace = "activity"

var (
	activityLabels = []string{"activity_level", "days_left"}

	activityPoints = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, activityNamespace, "points"),
		"Current activity level",
		append(labels, activityLabels...),
		nil,
	)
	activityMaintainPoints = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, activityNamespace, "maintain_points"),
		"Current activity level",
		append(labels, activityLabels...),
		nil,
	)
	activityGoal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, activityNamespace, "points_goal"),
		"Current activity level",
		append(labels, activityLabels...),
		nil,
	)
)

func (e *EgymExporter) describeActivityLevelMetrics(ch chan<- *prometheus.Desc) {
	ch <- activityPoints
	ch <- activityMaintainPoints
	ch <- activityGoal
}

func (e *EgymExporter) collectActivityLevelMetrics(ch chan<- prometheus.Metric) {
	data, err := e.client.GetActivityLevel()
	if err != nil {
		log.Error("Could not load activity level data", err)
		return
	}

	labelValues := []string{e.client.Username, data.Level, strconv.Itoa(data.DaysLeft)}

	ch <- prometheus.MustNewConstMetric(
		activityPoints,
		prometheus.GaugeValue,
		float64(data.Points),
		labelValues...,
	)
	ch <- prometheus.MustNewConstMetric(
		activityMaintainPoints,
		prometheus.GaugeValue,
		float64(data.MaintainPoints),
		labelValues...,
	)
	ch <- prometheus.MustNewConstMetric(
		activityGoal,
		prometheus.GaugeValue,
		float64(data.Goal),
		labelValues...,
	)
}
