package models

type ProblemDetails struct {
	Type               string          `json:"type,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
}
