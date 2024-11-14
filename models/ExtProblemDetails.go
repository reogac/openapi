/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtProblemDetails struct {
	NrfId              string          `json:"nrfId,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Title              string          `json:"title,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Type               string          `json:"type,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Detail             string          `json:"detail,omitempty"`
}
