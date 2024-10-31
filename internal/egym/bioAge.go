package egym

import (
	"encoding/json"
	"fmt"
)

func (c *EgymClient) GetBioAge() (*GetBioAgeResponse, error) {
	url := fmt.Sprintf("%s/analysis/api/v1.0/exercisers/%s/bioage", c.apiUrl, c.userId)
	body, err := c.fetch(url, 1)
	if err != nil {
		return nil, err
	}

	var response GetBioAgeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type GetBioAgeResponse struct {
	TotalDetails struct {
		TotalBioAge BioAgeDetail `json:"totalBioAge"`
	} `json:"totalDetails"`
	MuscleDetails struct {
		UpperBodyAge MuscleBodyAgeDetail `json:"upperBodyAge"`
		CoreAge      MuscleBodyAgeDetail `json:"coreAge"`
		LowerBodyAge MuscleBodyAgeDetail `json:"lowerBodyAge"`
		MuscleBioAge BioAgeDetail        `json:"muscleBioAge"`
	} `json:"muscleDetails"`
	MetabolicDetails struct {
		MetabolicAge BioAgeDetail `json:"metabolicAge"`
	} `json:"metabolicDetails"`
	CardioDetails struct {
		CardioAge BioAgeDetail `json:"cardioAge"`
		VO2Max    BioAgeDetail `json:"vo2max"`
	} `json:"cardioDetails"`
	FlexibilityDetails struct {
		FlexibilityAge BioAgeDetail `json:"flexibilityAge"`
	} `json:"flexibilityDetails"`
}

type BioAgeDetail struct {
	Value          float64 `json:"value"`
	Progress       string  `json:"progress"`
	PercentageDiff float64 `json:"percentageDiff"`
	AmountDiff     float64 `json:"amountDiff"`
	CreatedAt      string  `json:"createdAt"`
	Timezone       string  `json:"timezone"`
}

type MuscleBodyAgeDetail struct {
	Value          float64 `json:"value"`
	Progress       string  `json:"progress"`
	PercentageDiff float64 `json:"percentageDiff"`
	AmountDiff     float64 `json:"amountDiff"`
	CreatedAt      string  `json:"createdAt"`
	Timezone       string  `json:"timezone"`
	MusclesState   string  `json:"musclesState"`
}
