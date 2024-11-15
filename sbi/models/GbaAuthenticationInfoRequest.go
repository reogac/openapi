/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GbaAuthenticationInfoRequest struct {
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	AuthType              GbaAuthType            `json:"authType"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
}
