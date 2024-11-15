/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SorInfo struct {
	AckInd                  bool               `json:"ackInd"`
	Countersor              string             `json:"countersor,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
}
