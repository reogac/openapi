package models
type SessionRuleReport struct {
	 RuleIds	[]string	`json:"ruleIds"`
	 RuleStatus	RuleStatus	`json:"ruleStatus"`
	 SessRuleFailureCode	SessionRuleFailureCode	`json:"sessRuleFailureCode,omitempty"`
	 PolicyDecFailureReports	[]string	`json:"policyDecFailureReports,omitempty"`
}
