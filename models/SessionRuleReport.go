package models

type SessionRuleReport struct {
	PolicyDecFailureReports []string               `json:"policyDecFailureReports,omitempty"`
	RuleIds                 []string               `json:"ruleIds"`
	RuleStatus              RuleStatus             `json:"ruleStatus"`
	SessRuleFailureCode     SessionRuleFailureCode `json:"sessRuleFailureCode,omitempty"`
}
