package models

type ExtProblemDetails struct {
	Type               string          `json:"type,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Title              string          `json:"title,omitempty"`
	Detail             string          `json:"detail,omitempty"`
}
