package models
type RuleReport struct {
	 PccRuleIds	[]string	`json:"pccRuleIds"`
	 RuleStatus	RuleStatus	`json:"ruleStatus"`
	 ContVers	[]int	`json:"contVers,omitempty"`
	 FailureCode	FailureCode	`json:"failureCode,omitempty"`
	 FinUnitAct	FinalUnitAction	`json:"finUnitAct,omitempty"`
	 RanNasRelCauses	[]RanNasRelCause	`json:"ranNasRelCauses,omitempty"`
	 AltQosParamId	string	`json:"altQosParamId,omitempty"`
}
