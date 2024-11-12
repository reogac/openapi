package models

type PartialSuccessReport struct {
	InvalidPolicyDecs       []InvalidParam      `json:"invalidPolicyDecs,omitempty"`
	FailureCause            FailureCause        `json:"failureCause"`
	RuleReports             []RuleReport        `json:"ruleReports,omitempty"`
	SessRuleReports         []SessionRuleReport `json:"sessRuleReports,omitempty"`
	UeCampingRep            *UeCampingRep       `json:"ueCampingRep,omitempty"`
	PolicyDecFailureReports []string            `json:"policyDecFailureReports,omitempty"`
}
