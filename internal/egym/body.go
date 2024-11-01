package egym

import (
	"encoding/json"
	"fmt"
)

func (c *EgymClient) GetBodyMetrics() (*[]BodyMetric, error) {
	url := fmt.Sprintf("%s/measurements/api/v1.0/exercisers/%s/body/latest", c.apiUrl, c.userId)
	body, err := c.fetch(url, 1)
	if err != nil {
		return nil, err
	}

	var metrics []BodyMetric
	err = json.Unmarshal(body, &metrics)
	if err != nil {
		return nil, err
	}
	return &metrics, nil
}

type BodyMetric struct {
	Type        string  `json:"type"`
	Value       float64 `json:"value"`
	Source      string  `json:"source"`
	SourceLabel string  `json:"sourceLabel"`
}
