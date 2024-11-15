/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMtDataError struct {
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             *int            `json:"status,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	Type               string          `json:"type,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Detail             string          `json:"detail,omitempty"`
}
