package models

type UeInitiatedResourceRequest struct {
	Precedence   *int               `json:"precedence,omitempty"`
	PackFiltInfo []PacketFilterInfo `json:"packFiltInfo"`
	ReqQos       *RequestedQos      `json:"reqQos,omitempty"`
	PccRuleId    string             `json:"pccRuleId,omitempty"`
	RuleOp       RuleOperation      `json:"ruleOp"`
}
