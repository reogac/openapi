/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowSetupItem struct {
	Qfi                int               `json:"qfi"`
	QosRules           string            `json:"qosRules"`
	Ebi                *int              `json:"ebi,omitempty"`
	QosFlowDescription string            `json:"qosFlowDescription,omitempty"`
	QosFlowProfile     *QosFlowProfile   `json:"qosFlowProfile,omitempty"`
	AssociatedAnType   QosFlowAccessType `json:"associatedAnType,omitempty"`
	DefaultQosRuleInd  *bool             `json:"defaultQosRuleInd,omitempty"`
}
