package models

type ExtProblemDetails struct {
	Instance           string          `json:"instance,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Title              string          `json:"title,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Type               string          `json:"type,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Status             *int            `json:"status,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Cause              string          `json:"cause,omitempty"`
}
