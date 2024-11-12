package models

type UeInitiatedResourceRequest struct {
	PackFiltInfo []PacketFilterInfo `json:"packFiltInfo"`
	ReqQos       *RequestedQos      `json:"reqQos,omitempty"`
	PccRuleId    string             `json:"pccRuleId,omitempty"`
	RuleOp       RuleOperation      `json:"ruleOp"`
	Precedence   *int               `json:"precedence,omitempty"`
}
