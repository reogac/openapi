package models
type QosFlowSetupItem struct {
	 QosFlowDescription	string	`json:"qosFlowDescription,omitempty"`
	 QosFlowProfile	*QosFlowProfile	`json:"qosFlowProfile,omitempty"`
	 AssociatedAnType	QosFlowAccessType	`json:"associatedAnType,omitempty"`
	 DefaultQosRuleInd	*bool	`json:"defaultQosRuleInd,omitempty"`
	 Qfi	int	`json:"qfi"`
	 QosRules	string	`json:"qosRules"`
	 Ebi	*int	`json:"ebi,omitempty"`
}
