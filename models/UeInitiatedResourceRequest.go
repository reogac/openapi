package models
type UeInitiatedResourceRequest struct {
	 PccRuleId	string	`json:"pccRuleId,omitempty"`
	 RuleOp	RuleOperation	`json:"ruleOp"`
	 Precedence	*int	`json:"precedence,omitempty"`
	 PackFiltInfo	[]PacketFilterInfo	`json:"packFiltInfo"`
	 ReqQos	*RequestedQos	`json:"reqQos,omitempty"`
}
