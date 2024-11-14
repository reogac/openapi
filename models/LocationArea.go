package models
type LocationArea struct {
	 GeographicAreas	[]GeographicArea	`json:"geographicAreas,omitempty"`
	 CivicAddresses	[]CivicAddress	`json:"civicAddresses,omitempty"`
	 NwAreaInfo	*NetworkAreaInfo	`json:"nwAreaInfo,omitempty"`
	 UmtTime	*UmtTime	`json:"umtTime,omitempty"`
}
