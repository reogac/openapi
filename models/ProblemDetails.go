package models

type ProblemDetails struct {
	Instance           string          `json:"instance,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Type               string          `json:"type,omitempty"`
	Title              string          `json:"title,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
}
