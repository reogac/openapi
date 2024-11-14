package models
type ExtProblemDetails struct {
	 Detail	string	`json:"detail,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
	 Title	string	`json:"title,omitempty"`
	 Status	*int	`json:"status,omitempty"`
	 Instance	string	`json:"instance,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 RemoteError	*bool	`json:"remoteError,omitempty"`
}
