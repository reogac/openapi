package models
type DispersionRequirement struct {
	 Order	MatchingDirection	`json:"order,omitempty"`
	 DisperType	DispersionType	`json:"disperType"`
	 ClassCriters	[]ClassCriterion	`json:"classCriters,omitempty"`
	 RankCriters	[]RankingCriterion	`json:"rankCriters,omitempty"`
	 DispOrderCriter	DispersionOrderingCriterion	`json:"dispOrderCriter,omitempty"`
}
