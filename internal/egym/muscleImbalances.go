package egym

import (
	"encoding/json"
	"fmt"
)

func (c *EgymClient) GetMuscleImbalances() (*[]MuscleImbalance, error) {
	url := fmt.Sprintf("%s/analysis/api/v1.0/exercisers/%s/muscleimbalances", c.apiUrl, c.userId)
	body, err := c.fetch(url, 1)
	if err != nil {
		return nil, err
	}

	var response GetMuscleImbalancesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response.MuscleImbalances, nil
}

type GetMuscleImbalancesResponse struct {
	MuscleImbalances []MuscleImbalance `json:"muscleImbalances"`
}

type MuscleImbalance struct {
	CalculatedAt              string  `json:"calculatedAt"`
	Timezone                  string  `json:"timezone"`
	AgonistMuscle             string  `json:"agonistMuscle"`
	AntagonistMuscle          string  `json:"antagonistMuscle"`
	Position                  float64 `json:"position"`
	OptimalRangeStartPosition float64 `json:"optimalRangeStartPosition"`
	OptimalRangeEndPosition   float64 `json:"optimalRangeEndPosition"`
	RangeSize                 float64 `json:"rangeSize"`
	ImageUrl                  string  `json:"imageUrl"`
	BodyRegion                string  `json:"bodyRegion"`
	AgonistStrengthValue      float64 `json:"agonistStrengthValue"`
	AntagonistStrengthValue   float64 `json:"antagonistStrengthValue"`
	AgonistExerciseCode       string  `json:"agonistExerciseCode"`
	AntagonistExerciseCode    string  `json:"antagonistExerciseCode"`
	AgonistRatio              float64 `json:"agonistRatio"`
	AntagonistRatio           float64 `json:"antagonistRatio"`
}
