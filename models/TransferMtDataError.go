/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMtDataError struct {
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Type               string          `json:"type,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Status             *int            `json:"status,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Title              string          `json:"title,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Detail             string          `json:"detail,omitempty"`
}
