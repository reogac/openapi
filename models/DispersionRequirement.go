package models
type DispersionRequirement struct {
	 DispOrderCriter	DispersionOrderingCriterion	`json:"dispOrderCriter,omitempty"`
	 Order	MatchingDirection	`json:"order,omitempty"`
	 DisperType	DispersionType	`json:"disperType"`
	 ClassCriters	[]ClassCriterion	`json:"classCriters,omitempty"`
	 RankCriters	[]RankingCriterion	`json:"rankCriters,omitempty"`
}
