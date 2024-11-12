package models

type LocationArea struct {
	CivicAddresses  []CivicAddress   `json:"civicAddresses,omitempty"`
	NwAreaInfo      *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	UmtTime         *UmtTime         `json:"umtTime,omitempty"`
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`
}
