package models

type MatchingOperator string

// Define constant values for MatchingOperator
const (
	MATCHINGOPERATOR_FULL_MATCH     MatchingOperator = "FULL_MATCH"
	MATCHINGOPERATOR_MATCH_ALL      MatchingOperator = "MATCH_ALL"
	MATCHINGOPERATOR_STARTS_WITH    MatchingOperator = "STARTS_WITH"
	MATCHINGOPERATOR_NOT_START_WITH MatchingOperator = "NOT_START_WITH"
	MATCHINGOPERATOR_ENDS_WITH      MatchingOperator = "ENDS_WITH"
	MATCHINGOPERATOR_NOT_END_WITH   MatchingOperator = "NOT_END_WITH"
	MATCHINGOPERATOR_CONTAINS       MatchingOperator = "CONTAINS"
	MATCHINGOPERATOR_NOT_CONTAIN    MatchingOperator = "NOT_CONTAIN"
)