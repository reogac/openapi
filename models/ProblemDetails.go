package models
type ProblemDetails struct {
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 Title	string	`json:"title,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
	 Status	*int	`json:"status,omitempty"`
	 Detail	string	`json:"detail,omitempty"`
	 Instance	string	`json:"instance,omitempty"`
}
