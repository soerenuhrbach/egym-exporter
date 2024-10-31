package egym

import (
	"encoding/json"
	"fmt"
)

func (c *EgymClient) GetActivityLevel() (*GetActivityLevelsResponse, error) {
	url := fmt.Sprintf("%s/analysis/api/v1.0/exercisers/%s/activitylevels", c.apiUrl, c.userId)
	body, err := c.fetch(url, 1)
	if err != nil {
		return nil, err
	}

	var response GetActivityLevelsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type GetActivityLevelsResponse struct {
	Points         int    `json:"points"`
	DaysLeft       int    `json:"daysLeft"`
	Level          string `json:"level"`
	Goal           int    `json:"goal"`
	MaintainPoints int    `json:"maintainPoints"`
}
