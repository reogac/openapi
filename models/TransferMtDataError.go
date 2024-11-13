package models

type TransferMtDataError struct {
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Title              string          `json:"title,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Type               string          `json:"type,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
}
