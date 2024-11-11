package models

type DispersionRequirement struct {
	DisperType      string             `json:"disperType"`
	ClassCriters    []ClassCriterion   `json:"classCriters,omitempty"`
	RankCriters     []RankingCriterion `json:"rankCriters,omitempty"`
	DispOrderCriter string             `json:"dispOrderCriter,omitempty"`
	Order           string             `json:"order,omitempty"`
}
