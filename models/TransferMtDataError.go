package models

type TransferMtDataError struct {
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Title              string          `json:"title,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Type               string          `json:"type,omitempty"`
	Cause              string          `json:"cause,omitempty"`
}
