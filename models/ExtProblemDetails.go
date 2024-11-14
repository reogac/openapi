package models
type ExtProblemDetails struct {
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 RemoteError	*bool	`json:"remoteError,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 Status	*int	`json:"status,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
	 Detail	string	`json:"detail,omitempty"`
	 Instance	string	`json:"instance,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 Title	string	`json:"title,omitempty"`
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
}
