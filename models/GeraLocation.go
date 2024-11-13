package models

type GeraLocation struct {
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
}
