package models

type LocationArea struct {
	UmtTime         *UmtTime         `json:"umtTime,omitempty"`
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`
	CivicAddresses  []CivicAddress   `json:"civicAddresses,omitempty"`
	NwAreaInfo      *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
}
