package models

type UpuInfo struct {
	CounterUpu              string    `json:"counterUpu,omitempty"`
	ProvisioningTime        string    `json:"provisioningTime"`
	UpuTransparentContainer string    `json:"upuTransparentContainer,omitempty"`
	UpuDataList             []UpuData `json:"upuDataList,omitempty"`
	UpuRegInd               *bool     `json:"upuRegInd,omitempty"`
	UpuAckInd               *bool     `json:"upuAckInd,omitempty"`
	UpuMacIausf             string    `json:"upuMacIausf,omitempty"`
}
