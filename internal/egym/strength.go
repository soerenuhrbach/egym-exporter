package egym

import (
	"encoding/json"
	"fmt"
)

var (
	exerciseMuscleMap = map[string]string{
		"994":  "QUADRICEPS",
		"997":  "HAMSTRING",
		"996":  "LOWER_BACK",
		"995":  "ABS",
		"1000": "LATS",
		"1011": "SHOULDER",
		"999":  "UPPER_BACK",
		"998":  "CHEST",
		"1003": "OUTER_HIPS",
		"1004": "INNER_HIPS",
		"1001": "GLUTEUS",
	}
)

func (c *EgymClient) GetStrengthMetrics() (*[]StrengthMetric, error) {
	url := fmt.Sprintf("%s/measurements/api/v1.0/exercisers/%s/strength/latest", c.apiUrl, c.userId)
	body, err := c.fetch(url, 1)
	if err != nil {
		return nil, err
	}

	var response GetStrengthMetricsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response.StrengthMeasurements, nil
}

func (c *EgymClient) GetMuscleFromStrengthExercise(exerciseCode string) string {
	return exerciseMuscleMap[exerciseCode]
}

type GetStrengthMetricsResponse struct {
	StrengthMeasurements []StrengthMetric `json:"strengthMeasurements"`
}

type StrengthMetric struct {
	CreatedAt   string `json:"createdAt"`
	Timezone    string `json:"timezone"`
	Source      string `json:"source"`
	SourceLabel string `json:"sourceLabel"`
	BodyRegion  string `json:"bodyRegion"`
	Exercise    struct {
		Code  string `json:"code"`
		Label string `json:"label"`
	} `json:"exercise"`
	Strength struct {
		Value          float64 `json:"value"`
		Progress       string  `json:"progress"`
		PercentageDiff float64 `json:"percentageDiff"`
		AmountDiff     float64 `json:"amountDiff"`
		CreatedAt      string  `json:"createdAt"`
		Timezone       string  `json:"timezone"`
	} `json:"strength"`
}
