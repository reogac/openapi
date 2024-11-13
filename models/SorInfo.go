package models

type SorInfo struct {
	AckInd                  bool               `json:"ackInd"`
	Countersor              string             `json:"countersor,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
}
