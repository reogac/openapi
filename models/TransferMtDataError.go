package models

type TransferMtDataError struct {
	Type               string          `json:"type,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Title              string          `json:"title,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Status             *int            `json:"status,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
}
