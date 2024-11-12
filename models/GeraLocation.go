package models

type GeraLocation struct {
	MscNumber                string          `json:"mscNumber,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
}
