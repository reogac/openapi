package models
type ProblemDetails struct {
	 Title	string	`json:"title,omitempty"`
	 Detail	string	`json:"detail,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 Status	int	`json:"status"`
	 Instance	string	`json:"instance,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
}
