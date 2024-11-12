package models
type FqdnPatternMatchingRule struct {
	 Regex	string	`json:"regex,omitempty"`
	 StringMatchingRule	*StringMatchingRule	`json:"stringMatchingRule,omitempty"`
}
