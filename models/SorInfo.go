package models

type SorInfo struct {
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	Countersor              string             `json:"countersor,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
	AckInd                  bool               `json:"ackInd"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
}
