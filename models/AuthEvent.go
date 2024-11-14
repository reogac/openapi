package models
type AuthEvent struct {
	 NfSetId	string	`json:"nfSetId,omitempty"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 Success	bool	`json:"success"`
	 ServingNetworkName	string	`json:"servingNetworkName"`
	 AuthRemovalInd	*bool	`json:"authRemovalInd,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 TimeStamp	string	`json:"timeStamp"`
	 AuthType	AuthType	`json:"authType"`
}
