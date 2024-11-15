/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PartialSuccessReport struct {
	UeCampingRep            *UeCampingRep       `json:"ueCampingRep,omitempty"`
	PolicyDecFailureReports []string            `json:"policyDecFailureReports,omitempty"`
	InvalidPolicyDecs       []InvalidParam      `json:"invalidPolicyDecs,omitempty"`
	FailureCause            FailureCause        `json:"failureCause"`
	RuleReports             []RuleReport        `json:"ruleReports,omitempty"`
	SessRuleReports         []SessionRuleReport `json:"sessRuleReports,omitempty"`
}
