package models
type DnnInfo struct {
	 DnnBarred	*bool	`json:"dnnBarred,omitempty"`
	 InvokeNefInd	*bool	`json:"invokeNefInd,omitempty"`
	 SmfList	[]string	`json:"smfList,omitempty"`
	 SameSmfInd	*bool	`json:"sameSmfInd,omitempty"`
	 Dnn	string	`json:"dnn"`
	 DefaultDnnIndicator	*bool	`json:"defaultDnnIndicator,omitempty"`
	 LboRoamingAllowed	*bool	`json:"lboRoamingAllowed,omitempty"`
	 IwkEpsInd	*bool	`json:"iwkEpsInd,omitempty"`
}
