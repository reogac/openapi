package models
type QosFlowSetupItem struct {
	 QosFlowProfile	*QosFlowProfile	`json:"qosFlowProfile,omitempty"`
	 AssociatedAnType	QosFlowAccessType	`json:"associatedAnType,omitempty"`
	 DefaultQosRuleInd	*bool	`json:"defaultQosRuleInd,omitempty"`
	 Qfi	int	`json:"qfi"`
	 QosRules	string	`json:"qosRules"`
	 Ebi	*int	`json:"ebi,omitempty"`
	 QosFlowDescription	string	`json:"qosFlowDescription,omitempty"`
}
