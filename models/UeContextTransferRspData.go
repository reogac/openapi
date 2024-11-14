/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextTransferRspData struct {
	SupportedFeatures          string         `json:"supportedFeatures,omitempty"`
	UeContext                  UeContext      `json:"ueContext"`
	UeRadioCapability          *N2InfoContent `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent `json:"ueRadioCapabilityForPaging,omitempty"`
	UeNbiotRadioCapability     *N2InfoContent `json:"ueNbiotRadioCapability,omitempty"`
}
