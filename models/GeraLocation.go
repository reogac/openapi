package models

type GeraLocation struct {
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
}
