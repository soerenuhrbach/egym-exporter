package egym

import (
	"encoding/json"
	"fmt"
)

func (c *EgymClient) GetFlexibilityMetrics() (*[]FlexibilityMetric, error) {
	url := fmt.Sprintf("%s/measurements/api/v1.0/exercisers/%s/flexibility/latest", c.apiUrl, c.userId)
	body, err := c.fetch(url, 1)
	if err != nil {
		return nil, err
	}

	var response []FlexibilityMetric
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type FlexibilityMetric struct {
	Type                string  `json:"type"`
	Value               float64 `json:"value"`
	CreatedAt           string  `json:"createdAt"`
	Source              string  `json:"source"`
	SourceLabel         string  `json:"sourceLabel"`
	ValueInterpretation string  `json:"valueInterpretation"`
}
