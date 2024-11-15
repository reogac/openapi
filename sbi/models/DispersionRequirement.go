/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionRequirement struct {
	RankCriters     []RankingCriterion          `json:"rankCriters,omitempty"`
	DispOrderCriter DispersionOrderingCriterion `json:"dispOrderCriter,omitempty"`
	Order           MatchingDirection           `json:"order,omitempty"`
	DisperType      DispersionType              `json:"disperType"`
	ClassCriters    []ClassCriterion            `json:"classCriters,omitempty"`
}
