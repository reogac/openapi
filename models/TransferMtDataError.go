package models

type TransferMtDataError struct {
	Detail             string          `json:"detail,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Type               string          `json:"type,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             int             `json:"status"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	Instance           string          `json:"instance,omitempty"`
}
