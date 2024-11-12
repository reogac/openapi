package models

type ExtProblemDetails struct {
	Detail             string          `json:"detail,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	Status             *int            `json:"status,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Type               string          `json:"type,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Title              string          `json:"title,omitempty"`
	Instance           string          `json:"instance,omitempty"`
}
