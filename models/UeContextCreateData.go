/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreateData struct {
	UeContext                  UeContext         `json:"ueContext"`
	TargetId                   NgRanTargetId     `json:"targetId"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
}
