package models

type LocationArea struct {
	NwAreaInfo      *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	UmtTime         *UmtTime         `json:"umtTime,omitempty"`
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`
	CivicAddresses  []CivicAddress   `json:"civicAddresses,omitempty"`
}
