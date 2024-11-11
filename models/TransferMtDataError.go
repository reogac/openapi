package models

type TransferMtDataError struct {
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Type               string          `json:"type,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             *int            `json:"status,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
}
