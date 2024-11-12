package models

type DispersionRequirement struct {
	DisperType      DispersionType              `json:"disperType"`
	ClassCriters    []ClassCriterion            `json:"classCriters,omitempty"`
	RankCriters     []RankingCriterion          `json:"rankCriters,omitempty"`
	DispOrderCriter DispersionOrderingCriterion `json:"dispOrderCriter,omitempty"`
	Order           MatchingDirection           `json:"order,omitempty"`
}
