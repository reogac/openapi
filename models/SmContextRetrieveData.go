package models

type SmContextRetrieveData struct {
	NotToTransferEbiList []int            `json:"notToTransferEbiList,omitempty"`
	RanUnchangedInd      *bool            `json:"ranUnchangedInd,omitempty"`
	TargetMmeCap         *MmeCapabilities `json:"targetMmeCap,omitempty"`
	SmContextType        string           `json:"smContextType,omitempty"`
	ServingNetwork       *PlmnId          `json:"servingNetwork,omitempty"`
}
