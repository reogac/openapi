package models

type SorInfo struct {
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
	Countersor              string             `json:"countersor,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	AckInd                  bool               `json:"ackInd"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
}
