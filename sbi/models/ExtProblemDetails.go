/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtProblemDetails struct {
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Status             *int            `json:"status,omitempty"`
	Title              string          `json:"title,omitempty"`
	Type               string          `json:"type,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
}
