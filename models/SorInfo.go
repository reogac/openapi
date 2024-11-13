package models

type SorInfo struct {
	AckInd                  bool               `json:"ackInd"`
	Countersor              string             `json:"countersor,omitempty"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
}
