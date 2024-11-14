package models
type GeoServiceArea struct {
	 CivicAddressList	[]CivicAddress	`json:"civicAddressList,omitempty"`
	 GeographicAreaList	[]GeographicArea	`json:"geographicAreaList,omitempty"`
}
