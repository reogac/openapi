package models

type AcknowledgeInfo struct {
	UpuTransparentContainer string `json:"upuTransparentContainer,omitempty"`
	SorMacIue               string `json:"sorMacIue,omitempty"`
	UpuMacIue               string `json:"upuMacIue,omitempty"`
	ProvisioningTime        string `json:"provisioningTime"`
	SorTransparentContainer string `json:"sorTransparentContainer,omitempty"`
	UeNotReachable          *bool  `json:"ueNotReachable,omitempty"`
}
