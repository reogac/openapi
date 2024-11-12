package models

type TransferMtDataError struct {
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Type               string          `json:"type,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Title              string          `json:"title,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	Detail             string          `json:"detail,omitempty"`
}
