package models

type GeraLocation struct {
	LocationNumber           string          `json:"locationNumber,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
}
