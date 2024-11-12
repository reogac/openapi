package models

type TransferMtDataError struct {
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Type               string          `json:"type,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Title              string          `json:"title,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Status             *int            `json:"status,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
}
