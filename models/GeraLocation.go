package models

type GeraLocation struct {
	MscNumber                string          `json:"mscNumber,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
}
