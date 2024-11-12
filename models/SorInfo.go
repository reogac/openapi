package models

type SorInfo struct {
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	AckInd                  bool               `json:"ackInd"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	Countersor              string             `json:"countersor,omitempty"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
}
