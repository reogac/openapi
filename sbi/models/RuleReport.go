/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RuleReport struct {
	RuleStatus      RuleStatus       `json:"ruleStatus"`
	ContVers        []int            `json:"contVers,omitempty"`
	FailureCode     FailureCode      `json:"failureCode,omitempty"`
	FinUnitAct      FinalUnitAction  `json:"finUnitAct,omitempty"`
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`
	AltQosParamId   string           `json:"altQosParamId,omitempty"`
	PccRuleIds      []string         `json:"pccRuleIds"`
}
