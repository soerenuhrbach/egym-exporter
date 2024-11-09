package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/soerenuhrbach/egym-exporter/internal/egym"

	"time"

	log "github.com/sirupsen/logrus"
)

const exerciseNamespace = "exercise"

var (
	exerciseLabels = []string{"exercise", "code", "completed_at", "source", "source_label", "unit", "workout"}

	exerciseActivityPoints = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exerciseNamespace, "activity_points"),
		"Activity points of the exercise",
		append(labels, exerciseLabels...),
		nil,
	)
	exerciseDistance = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exerciseNamespace, "distance"),
		"Distance of the exercise",
		append(labels, exerciseLabels...),
		nil,
	)
	exerciseDuration = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exerciseNamespace, "duration"),
		"Training duration of the exercise",
		append(labels, exerciseLabels...),
		nil,
	)
	exerciseCalories = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exerciseNamespace, "calories"),
		"Burned calories with the exercise",
		append(labels, exerciseLabels...),
		nil,
	)
	exerciseAverageSpeed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exerciseNamespace, "average_speed"),
		"Average speed of the exercise",
		append(labels, exerciseLabels...),
		nil,
	)
	exerciseSets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exerciseNamespace, "sets"),
		"Number of sets of the exercise",
		append(labels, exerciseLabels...),
		nil,
	)
	exerciseReps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exerciseNamespace, "reps"),
		"Total number of reps over all sets of the exercise",
		append(labels, exerciseLabels...),
		nil,
	)
	exerciseWeightPerRep = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exerciseNamespace, "weight_per_rep"),
		"Average weight per rep of the exercise",
		append(labels, exerciseLabels...),
		nil,
	)
	exerciseTotalWeight = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exerciseNamespace, "weight_total"),
		"Total weight across all reps of all sets of the exercise",
		append(labels, exerciseLabels...),
		nil,
	)
)

func (c *EgymExporter) describeExerciseMetrics(ch chan<- *prometheus.Desc) {
	ch <- exerciseActivityPoints
}

func (c *EgymExporter) collectExerciseMetrics(ch chan<- prometheus.Metric) {
	now := time.Now().UTC()
	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		int(0),
		int(0),
		int(0),
		int(0),
		now.Location(),
	)
	endOfDay := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		int(23),
		int(59),
		int(59),
		int(59),
		now.Location(),
	)

	workouts, err := c.client.GetWorkoutsInPeriod(today, endOfDay)
	if err != nil {
		log.Error("could not retrieve exercise data!", err)
		return
	}

	for _, workout := range *workouts {
		for _, exercise := range workout.Exercises {

			if exercise.Attributes.ActivityPoints != nil && exercise.Attributes.ActivityPoints.Value > 0 {
				ch <- c.createExerciseMetric(exerciseActivityPoints, workout, exercise, exercise.Attributes.ActivityPoints.Value, exercise.Attributes.ActivityPoints.Unit)
			}
			if exercise.Attributes.Distance != nil && exercise.Attributes.Distance.Value > 0 {
				ch <- c.createExerciseMetric(exerciseDistance, workout, exercise, exercise.Attributes.Distance.Value, exercise.Attributes.Distance.Unit)
			}
			if exercise.Attributes.Duration != nil && exercise.Attributes.Duration.Value > 0 {
				ch <- c.createExerciseMetric(exerciseDuration, workout, exercise, exercise.Attributes.Duration.Value, exercise.Attributes.Duration.Unit)
			}
			if exercise.Attributes.Calories != nil && exercise.Attributes.Calories.Value > 0 {
				ch <- c.createExerciseMetric(exerciseCalories, workout, exercise, exercise.Attributes.Calories.Value, exercise.Attributes.Calories.Unit)
			}
			if exercise.Attributes.AverageSpeed != nil && exercise.Attributes.AverageSpeed.Value > 0 {
				ch <- c.createExerciseMetric(exerciseAverageSpeed, workout, exercise, exercise.Attributes.AverageSpeed.Value, exercise.Attributes.AverageSpeed.Unit)
			}

			if exercise.Attributes.Sets != nil && len(exercise.Attributes.Sets) > 0 {
				sets := float64(len(exercise.Attributes.Sets))
				reps := float64(0)
				weight := float64(0)

				for _, set := range exercise.Attributes.Sets {
					if set.Reps != nil {
						reps += set.Reps.Value
					}

					if set.Weight != nil {
						weight += set.Weight.Value
					}
				}

				weight = weight / sets

				ch <- c.createExerciseMetric(exerciseSets, workout, exercise, sets, "")
				ch <- c.createExerciseMetric(exerciseReps, workout, exercise, reps, "")
				ch <- c.createExerciseMetric(exerciseWeightPerRep, workout, exercise, weight, "kg")
				ch <- c.createExerciseMetric(exerciseTotalWeight, workout, exercise, weight*reps, "kg")
			}
		}
	}
}

func (c *EgymExporter) createExerciseMetric(desc *prometheus.Desc, workout egym.Workout, exercise egym.Exercise, value float64, unit string) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		desc,
		prometheus.CounterValue,
		value,
		c.client.Username,
		exercise.Exercise.Label,
		exercise.Exercise.Code,
		exercise.CompletedAt,
		exercise.Source.Code,
		exercise.Source.Label,
		unit,
		workout.Code,
	)
}
