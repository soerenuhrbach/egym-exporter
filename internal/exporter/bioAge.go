package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const bioAgeNamespace = "bio_age"

var (
	totalBioAge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, bioAgeNamespace, "total"),
		"Total bio age",
		labels,
		nil,
	)
	musclesBioAge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, bioAgeNamespace, "muscles"),
		"Total muscle bio age",
		labels,
		nil,
	)
	upperBodyMusclesBioAge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, bioAgeNamespace, "muscles_upper_body"),
		"Upper body muscle bio age",
		labels,
		nil,
	)
	coreMusclesBioAge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, bioAgeNamespace, "muscles_core"),
		"Core body muscle bio age",
		labels,
		nil,
	)
	lowerBodyMusclesBioAge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, bioAgeNamespace, "muscles_lower_body"),
		"Lower body muscle bio age",
		labels,
		nil,
	)
	metabolicBioAge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, bioAgeNamespace, "metabolic"),
		"Metabolic bio age",
		labels,
		nil,
	)
	cardioBioAge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, bioAgeNamespace, "cardio"),
		"Cardio bio age",
		labels,
		nil,
	)
	flexibilityAge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, bioAgeNamespace, "flexibility"),
		"Flexibility bio age",
		labels,
		nil,
	)
)

func (c *EgymExporter) describeBioAgeMetrics(ch chan<- *prometheus.Desc) {
	ch <- totalBioAge
	ch <- musclesBioAge
	ch <- lowerBodyMusclesBioAge
	ch <- upperBodyMusclesBioAge
	ch <- coreMusclesBioAge
	ch <- metabolicBioAge
	ch <- cardioBioAge
	ch <- flexibilityAge
}

func (c *EgymExporter) collectBioAgeMetrics(ch chan<- prometheus.Metric) {
	bioAge, err := c.client.GetBioAge()
	if err != nil {
		log.Error("could not collect metrics for bioAge!", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		totalBioAge,
		prometheus.GaugeValue,
		bioAge.TotalDetails.TotalBioAge.Value,
		c.client.Username,
	)
	ch <- prometheus.MustNewConstMetric(
		musclesBioAge,
		prometheus.GaugeValue,
		bioAge.MuscleDetails.MuscleBioAge.Value,
		c.client.Username,
	)
	ch <- prometheus.MustNewConstMetric(
		lowerBodyMusclesBioAge,
		prometheus.GaugeValue,
		bioAge.MuscleDetails.LowerBodyAge.Value,
		c.client.Username,
	)
	ch <- prometheus.MustNewConstMetric(
		upperBodyMusclesBioAge,
		prometheus.GaugeValue,
		bioAge.MuscleDetails.UpperBodyAge.Value,
		c.client.Username,
	)
	ch <- prometheus.MustNewConstMetric(
		coreMusclesBioAge,
		prometheus.GaugeValue,
		bioAge.MuscleDetails.CoreAge.Value,
		c.client.Username,
	)
	ch <- prometheus.MustNewConstMetric(
		metabolicBioAge,
		prometheus.GaugeValue,
		bioAge.MetabolicDetails.MetabolicAge.Value,
		c.client.Username,
	)
	ch <- prometheus.MustNewConstMetric(
		cardioBioAge,
		prometheus.GaugeValue,
		bioAge.CardioDetails.CardioAge.Value,
		c.client.Username,
	)
	ch <- prometheus.MustNewConstMetric(
		flexibilityAge,
		prometheus.GaugeValue,
		bioAge.FlexibilityDetails.FlexibilityAge.Value,
		c.client.Username,
	)
}
