package models
type SessionRuleReport struct {
	 SessRuleFailureCode	SessionRuleFailureCode	`json:"sessRuleFailureCode,omitempty"`
	 PolicyDecFailureReports	[]string	`json:"policyDecFailureReports,omitempty"`
	 RuleIds	[]string	`json:"ruleIds"`
	 RuleStatus	RuleStatus	`json:"ruleStatus"`
}
