package models

type GeraLocation struct {
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
}
