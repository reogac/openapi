package models

type GeraLocation struct {
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
}
