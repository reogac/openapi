package models
type AuthEvent struct {
	 Success	bool	`json:"success"`
	 TimeStamp	string	`json:"timeStamp"`
	 AuthType	AuthType	`json:"authType"`
	 ServingNetworkName	string	`json:"servingNetworkName"`
	 NfSetId	string	`json:"nfSetId,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 AuthRemovalInd	*bool	`json:"authRemovalInd,omitempty"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
}
