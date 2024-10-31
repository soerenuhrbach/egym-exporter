package egym

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
		// BodyFat      BioAgeDetail `json:"bodyFat"`
		// BMI          BioAgeDetail `json:"bmi"`
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
