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

func (c *EgymExporter) Describe(ch chan<- *prometheus.Desc) {
	c.describeBioAgeMetrics(ch)
}

func (c *EgymExporter) Collect(ch chan<- prometheus.Metric) {
	c.collectBioAgeMetrics(ch)
}

func NewEgymExporter(client *egym.EgymClient) *EgymExporter {
	return &EgymExporter{
		client: client,
	}
}
