package models

type AccNetChId struct {
	AccNetChargId    string   `json:"accNetChargId,omitempty"`
	RefPccRuleIds    []string `json:"refPccRuleIds,omitempty"`
	SessionChScope   *bool    `json:"sessionChScope,omitempty"`
	AccNetChaIdValue *int     `json:"accNetChaIdValue,omitempty"`
}
