package models
type ProblemDetails struct {
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 Title	string	`json:"title,omitempty"`
	 Instance	string	`json:"instance,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 Status	*int	`json:"status,omitempty"`
	 Detail	string	`json:"detail,omitempty"`
}
