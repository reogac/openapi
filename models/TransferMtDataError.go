package models

type TransferMtDataError struct {
	Type               string          `json:"type,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Title              string          `json:"title,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
}
