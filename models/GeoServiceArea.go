package models

type GeoServiceArea struct {
	GeographicAreaList []GeographicArea `json:"geographicAreaList,omitempty"`
	CivicAddressList   []CivicAddress   `json:"civicAddressList,omitempty"`
}
