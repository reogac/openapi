package models

type TransferMtDataError struct {
	Instance           string          `json:"instance,omitempty"`
	Status             *int            `json:"status,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Title              string          `json:"title,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Type               string          `json:"type,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
}
