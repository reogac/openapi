package models

type TransferMtDataError struct {
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Status             int             `json:"status"`
	Instance           string          `json:"instance,omitempty"`
	Title              string          `json:"title,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Type               string          `json:"type,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	Detail             string          `json:"detail,omitempty"`
}
