package models

type SmContextRetrieveData struct {
	SmContextType        SmContextType    `json:"smContextType,omitempty"`
	ServingNetwork       *PlmnId          `json:"servingNetwork,omitempty"`
	NotToTransferEbiList []int            `json:"notToTransferEbiList,omitempty"`
	RanUnchangedInd      *bool            `json:"ranUnchangedInd,omitempty"`
	TargetMmeCap         *MmeCapabilities `json:"targetMmeCap,omitempty"`
}
