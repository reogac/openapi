/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMtDataError struct {
	Cause              string          `json:"cause,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Title              string          `json:"title,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	Type               string          `json:"type,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Status             *int            `json:"status,omitempty"`
}
