/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionRequirement struct {
	Order           MatchingDirection           `json:"order,omitempty"`
	DisperType      DispersionType              `json:"disperType"`
	ClassCriters    []ClassCriterion            `json:"classCriters,omitempty"`
	RankCriters     []RankingCriterion          `json:"rankCriters,omitempty"`
	DispOrderCriter DispersionOrderingCriterion `json:"dispOrderCriter,omitempty"`
}
