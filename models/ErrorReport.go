/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ErrorReport struct {
	SessRuleReports      []SessionRuleReport `json:"sessRuleReports,omitempty"`
	PolDecFailureReports []string            `json:"polDecFailureReports,omitempty"`
	InvalidPolicyDecs    []InvalidParam      `json:"invalidPolicyDecs,omitempty"`
	Error                *ProblemDetails     `json:"error,omitempty"`
	RuleReports          []RuleReport        `json:"ruleReports,omitempty"`
}
