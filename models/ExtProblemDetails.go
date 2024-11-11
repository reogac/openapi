package models

type ExtProblemDetails struct {
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Type               string          `json:"type,omitempty"`
}
