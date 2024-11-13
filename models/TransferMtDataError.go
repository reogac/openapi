package models

type TransferMtDataError struct {
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Type               string          `json:"type,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Title              string          `json:"title,omitempty"`
}
