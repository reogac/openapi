package models

type ExternalMbsServiceArea struct {
	GeographicAreaList []GeographicArea `json:"geographicAreaList,omitempty"`
	CivicAddressList   []CivicAddress   `json:"civicAddressList,omitempty"`
}
