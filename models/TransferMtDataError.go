package models
type TransferMtDataError struct {
	 Status	*int	`json:"status,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 Title	string	`json:"title,omitempty"`
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 RemoteError	*bool	`json:"remoteError,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 MaxWaitingTime	*int	`json:"maxWaitingTime,omitempty"`
	 Instance	string	`json:"instance,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
	 Detail	string	`json:"detail,omitempty"`
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
}
