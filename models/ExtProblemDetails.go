package models

type ExtProblemDetails struct {
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	Type               string          `json:"type,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Title              string          `json:"title,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Status             *int            `json:"status,omitempty"`
}
