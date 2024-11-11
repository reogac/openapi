package models
type SmContextRetrieveData struct {
	 RanUnchangedInd	*bool	`json:"ranUnchangedInd,omitempty"`
	 TargetMmeCap	*MmeCapabilities	`json:"targetMmeCap,omitempty"`
	 SmContextType	string	`json:"smContextType,omitempty"`
	 ServingNetwork	*PlmnId	`json:"servingNetwork,omitempty"`
	 NotToTransferEbiList	[]int	`json:"notToTransferEbiList,omitempty"`
}
