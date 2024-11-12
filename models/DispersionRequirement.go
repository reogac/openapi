package models

type DispersionRequirement struct {
	ClassCriters    []ClassCriterion            `json:"classCriters,omitempty"`
	RankCriters     []RankingCriterion          `json:"rankCriters,omitempty"`
	DispOrderCriter DispersionOrderingCriterion `json:"dispOrderCriter,omitempty"`
	Order           MatchingDirection           `json:"order,omitempty"`
	DisperType      DispersionType              `json:"disperType"`
}
