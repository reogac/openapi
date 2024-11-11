package models

type ExtProblemDetails struct {
	Type               string          `json:"type,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             int             `json:"status"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
}
