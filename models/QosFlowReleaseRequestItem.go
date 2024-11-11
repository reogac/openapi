package models

type QosFlowReleaseRequestItem struct {
	QosFlowDescription string `json:"qosFlowDescription,omitempty"`
	Qfi                int    `json:"qfi"`
	QosRules           string `json:"qosRules,omitempty"`
}
