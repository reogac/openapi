package models
type SorInfo struct {
	 AckInd	bool	`json:"ackInd"`
	 SorMacIausf	string	`json:"sorMacIausf,omitempty"`
	 ProvisioningTime	string	`json:"provisioningTime"`
	 StoreSorCmciInMe	*bool	`json:"storeSorCmciInMe,omitempty"`
	 SteeringContainer	*SteeringContainer	`json:"steeringContainer,omitempty"`
	 Countersor	string	`json:"countersor,omitempty"`
	 SorTransparentContainer	string	`json:"sorTransparentContainer,omitempty"`
	 SorCmci	string	`json:"sorCmci,omitempty"`
	 UsimSupportOfSorCmci	*bool	`json:"usimSupportOfSorCmci,omitempty"`
}
