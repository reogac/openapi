package models

type FqdnPatternMatchingRule struct {
	StringMatchingRule *StringMatchingRule `json:"stringMatchingRule,omitempty"`
	Regex              string              `json:"regex,omitempty"`
}
