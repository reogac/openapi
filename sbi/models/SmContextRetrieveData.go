/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextRetrieveData struct {
	SmContextType        SmContextType    `json:"smContextType,omitempty"`
	ServingNetwork       *PlmnId          `json:"servingNetwork,omitempty"`
	NotToTransferEbiList []int            `json:"notToTransferEbiList,omitempty"`
	RanUnchangedInd      *bool            `json:"ranUnchangedInd,omitempty"`
	TargetMmeCap         *MmeCapabilities `json:"targetMmeCap,omitempty"`
}
