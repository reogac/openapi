package models

type RequestedRuleData struct {
	RefPccRuleIds []string `json:"refPccRuleIds"`
	ReqData       []string `json:"reqData"`
}
