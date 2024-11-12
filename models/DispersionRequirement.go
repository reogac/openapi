package models
type DispersionRequirement struct {
	 RankCriters	[]RankingCriterion	`json:"rankCriters,omitempty"`
	 DispOrderCriter	DispersionOrderingCriterion	`json:"dispOrderCriter,omitempty"`
	 Order	MatchingDirection	`json:"order,omitempty"`
	 DisperType	DispersionType	`json:"disperType"`
	 ClassCriters	[]ClassCriterion	`json:"classCriters,omitempty"`
}
