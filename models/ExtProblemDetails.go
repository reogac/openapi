package models

type ExtProblemDetails struct {
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Type               string          `json:"type,omitempty"`
	Title              string          `json:"title,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
}
