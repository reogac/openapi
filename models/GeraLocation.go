package models

type GeraLocation struct {
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
}
