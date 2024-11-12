package models

type SorInfo struct {
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	AckInd                  bool               `json:"ackInd"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	Countersor              string             `json:"countersor,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
}
