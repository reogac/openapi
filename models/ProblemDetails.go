package models
type ProblemDetails struct {
	 Title	string	`json:"title,omitempty"`
	 Instance	string	`json:"instance,omitempty"`
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 Status	*int	`json:"status,omitempty"`
	 Detail	string	`json:"detail,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
