package models
type UeInitiatedResourceRequest struct {
	 RuleOp	RuleOperation	`json:"ruleOp"`
	 Precedence	*int	`json:"precedence,omitempty"`
	 PackFiltInfo	[]PacketFilterInfo	`json:"packFiltInfo"`
	 ReqQos	*RequestedQos	`json:"reqQos,omitempty"`
	 PccRuleId	string	`json:"pccRuleId,omitempty"`
}
