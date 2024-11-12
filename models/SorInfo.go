package models

type SorInfo struct {
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	AckInd                  bool               `json:"ackInd"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
	Countersor              string             `json:"countersor,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
}
