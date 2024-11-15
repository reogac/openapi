/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeInitiatedResourceRequest struct {
	RuleOp       RuleOperation      `json:"ruleOp"`
	Precedence   *int               `json:"precedence,omitempty"`
	PackFiltInfo []PacketFilterInfo `json:"packFiltInfo"`
	ReqQos       *RequestedQos      `json:"reqQos,omitempty"`
	PccRuleId    string             `json:"pccRuleId,omitempty"`
}
