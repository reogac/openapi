package models
type CagData struct {
	 CagInfos	map[string]CagInfo	`json:"cagInfos"`
	 ProvisioningTime	string	`json:"provisioningTime,omitempty"`
}
