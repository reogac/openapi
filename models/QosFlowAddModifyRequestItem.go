package models
type QosFlowAddModifyRequestItem struct {
	 AssociatedAnType	QosFlowAccessType	`json:"associatedAnType,omitempty"`
	 Qfi	int	`json:"qfi"`
	 Ebi	*int	`json:"ebi,omitempty"`
	 QosRules	string	`json:"qosRules,omitempty"`
	 QosFlowDescription	string	`json:"qosFlowDescription,omitempty"`
	 QosFlowProfile	*QosFlowProfile	`json:"qosFlowProfile,omitempty"`
}
