package models

type StringMatchingCondition struct {
	MatchingString   string           `json:"matchingString,omitempty"`
	MatchingOperator MatchingOperator `json:"matchingOperator"`
}
