package models
type GeraLocation struct {
	 Cgi	*CellGlobalId	`json:"cgi,omitempty"`
	 Rai	*RoutingAreaId	`json:"rai,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
	 LocationNumber	string	`json:"locationNumber,omitempty"`
	 Sai	*ServiceAreaId	`json:"sai,omitempty"`
	 Lai	*LocationAreaId	`json:"lai,omitempty"`
	 VlrNumber	string	`json:"vlrNumber,omitempty"`
	 MscNumber	string	`json:"mscNumber,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
}
