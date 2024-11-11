package models

type ProblemDetails struct {
	Status             int             `json:"status"`
	Instance           string          `json:"instance,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Type               string          `json:"type,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Title              string          `json:"title,omitempty"`
}
