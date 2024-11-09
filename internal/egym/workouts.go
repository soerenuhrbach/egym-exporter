package egym

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *EgymClient) GetWorkoutsInPeriod(startDate, endDate time.Time) (*[]Workout, error) {
	url := fmt.Sprintf(
		"%s/workouts/api/workouts/v2.3/exercisers/%s/workouts?completedAfter=%s&completedBefore=%s",
		c.brandApiUrl,
		c.userId,
		startDate.Format(time.RFC3339),
		endDate.Format(time.RFC3339),
	)
	body, err := c.fetch(url, 1)
	if err != nil {
		return nil, err
	}

	var response GetWorkoutsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response.Workouts, nil
}

type GetWorkoutsResponse struct {
	Workouts []Workout `json:"workouts"`
}

type Workout struct {
	Code                 string     `json:"code"`
	Exercises            []Exercise `json:"exercises"`
	CreatedAt            string     `json:"createdAt"`
	UpdatedAt            string     `json:"updatedAt"`
	CompletedAt          string     `json:"completedAt"`
	Timezone             string     `json:"timezone"`
	WorkoutPlanCode      string     `json:"workoutPlanCode"`
	WorkoutPlanLabel     string     `json:"workoutPlanLabel"`
	WorkoutPlanImageUrl  string     `json:"workoutPlanImageUrl"`
	WorkoutPlanGroupType string     `json:"workoutPlanGroupType"`
}

type Exercise struct {
	Code         string `json:"code"`
	ExerciseCode string `json:"exerciseCode"`
	LibraryCode  string `json:"libraryCode"`
	Source       struct {
		Label string `json:"label"`
		Code  string `json:"code"`
	} `json:"source"`
	Exercise struct {
		Label       string   `json:"label"`
		Code        string   `json:"code"`
		Icons       []string `json:"icons"`
		Videos      []string `json:"videos"`
		Previews    []string `json:"previews"`
		Description string   `json:"description"`
		Synonyms    []string `json:"synonyms"`
		Category    struct {
			Code  string `json:"code"`
			Label string `json:"label"`
		} `json:"category"`
		MachineBased bool `json:"machineBased"`
	} `json:"exercise"`
	Attributes struct {
		Distance       *ValueWithUnit `json:"distance"`
		Duration       *ValueWithUnit `json:"duration"`
		Calories       *ValueWithUnit `json:"calories"`
		ActivityPoints *ValueWithUnit `json:"activity_points"`
		AverageSpeed   *ValueWithUnit `json:"average_speed"`
		Sets           []struct {
			Reps     *ValueWithUnit `json:"reps"`
			Duration *ValueWithUnit `json:"duration"`
			Weight   *ValueWithUnit `json:"weight"`
		} `json:"sets_of_reps_and_weight_or_duration_and_weight"`
	} `json:"attributes"`
	Name        string `json:"name"`
	Editable    bool   `json:"editable"`
	Deletable   bool   `json:"deletable"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	CompletedAt string `json:"completedAt"`
	Timezone    string `json:"timezone"`
	Origin      string `json:"origin"`
}

type ValueWithUnit struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}
