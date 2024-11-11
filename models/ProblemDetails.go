package models
type ProblemDetails struct {
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 Status	*int	`json:"status,omitempty"`
	 Detail	string	`json:"detail,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 Title	string	`json:"title,omitempty"`
	 Instance	string	`json:"instance,omitempty"`
}
