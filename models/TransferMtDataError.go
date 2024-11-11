type TransferMtDataError struct {
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 Instance	string	`json:"instance,omitempty"`
	 MaxWaitingTime	*int	`json:"maxWaitingTime,omitempty"`
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 Title	string	`json:"title,omitempty"`
	 Status	int	`json:"status"`
	 Detail	string	`json:"detail,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
}
