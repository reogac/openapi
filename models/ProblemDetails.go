package models

type ProblemDetails struct {
	Type               string          `json:"type,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Title              string          `json:"title,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
}
