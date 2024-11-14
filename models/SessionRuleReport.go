/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionRuleReport struct {
	SessRuleFailureCode     SessionRuleFailureCode `json:"sessRuleFailureCode,omitempty"`
	PolicyDecFailureReports []string               `json:"policyDecFailureReports,omitempty"`
	RuleIds                 []string               `json:"ruleIds"`
	RuleStatus              RuleStatus             `json:"ruleStatus"`
}
