package models
type AccNetChId struct {
	 SessionChScope	*bool	`json:"sessionChScope,omitempty"`
	 AccNetChaIdValue	*int	`json:"accNetChaIdValue,omitempty"`
	 AccNetChargId	string	`json:"accNetChargId,omitempty"`
	 RefPccRuleIds	[]string	`json:"refPccRuleIds,omitempty"`
}
