package models

type ExternalMbsServiceArea struct {
	CivicAddressList   []CivicAddress   `json:"civicAddressList,omitempty"`
	GeographicAreaList []GeographicArea `json:"geographicAreaList,omitempty"`
}
