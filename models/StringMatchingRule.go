package models

type StringMatchingRule struct {
	StringMatchingConditions []StringMatchingCondition `json:"stringMatchingConditions,omitempty"`
}
