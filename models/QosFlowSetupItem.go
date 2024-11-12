package models

type QosFlowSetupItem struct {
	AssociatedAnType   string          `json:"associatedAnType,omitempty"`
	DefaultQosRuleInd  *bool           `json:"defaultQosRuleInd,omitempty"`
	Qfi                int             `json:"qfi"`
	QosRules           string          `json:"qosRules"`
	Ebi                *int            `json:"ebi,omitempty"`
	QosFlowDescription string          `json:"qosFlowDescription,omitempty"`
	QosFlowProfile     *QosFlowProfile `json:"qosFlowProfile,omitempty"`
}
