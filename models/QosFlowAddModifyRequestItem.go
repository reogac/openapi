package models

type QosFlowAddModifyRequestItem struct {
	QosFlowProfile     *QosFlowProfile `json:"qosFlowProfile,omitempty"`
	AssociatedAnType   string          `json:"associatedAnType,omitempty"`
	Qfi                int             `json:"qfi"`
	Ebi                *int            `json:"ebi,omitempty"`
	QosRules           string          `json:"qosRules,omitempty"`
	QosFlowDescription string          `json:"qosFlowDescription,omitempty"`
}
