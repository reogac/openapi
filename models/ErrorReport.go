package models

type ErrorReport struct {
	RuleReports          []RuleReport        `json:"ruleReports,omitempty"`
	SessRuleReports      []SessionRuleReport `json:"sessRuleReports,omitempty"`
	PolDecFailureReports []string            `json:"polDecFailureReports,omitempty"`
	InvalidPolicyDecs    []InvalidParam      `json:"invalidPolicyDecs,omitempty"`
	Error                *ProblemDetails     `json:"error,omitempty"`
}
