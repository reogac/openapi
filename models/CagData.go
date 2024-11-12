package models

type CagData struct {
	ProvisioningTime string             `json:"provisioningTime,omitempty"`
	CagInfos         map[string]CagInfo `json:"cagInfos"`
}
