package models
type SorInfo struct {
	 SteeringContainer	*SteeringContainer	`json:"steeringContainer,omitempty"`
	 AckInd	bool	`json:"ackInd"`
	 SorTransparentContainer	string	`json:"sorTransparentContainer,omitempty"`
	 SorMacIausf	string	`json:"sorMacIausf,omitempty"`
	 Countersor	string	`json:"countersor,omitempty"`
	 ProvisioningTime	string	`json:"provisioningTime"`
	 SorCmci	string	`json:"sorCmci,omitempty"`
	 StoreSorCmciInMe	*bool	`json:"storeSorCmciInMe,omitempty"`
	 UsimSupportOfSorCmci	*bool	`json:"usimSupportOfSorCmci,omitempty"`
}
