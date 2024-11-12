package models

type TransferMtDataError struct {
	Type               string          `json:"type,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Status             *int            `json:"status,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Title              string          `json:"title,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
}
