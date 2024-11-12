package models
type GeraLocation struct {
	 LocationNumber	string	`json:"locationNumber,omitempty"`
	 Cgi	*CellGlobalId	`json:"cgi,omitempty"`
	 Sai	*ServiceAreaId	`json:"sai,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
	 Lai	*LocationAreaId	`json:"lai,omitempty"`
	 Rai	*RoutingAreaId	`json:"rai,omitempty"`
	 VlrNumber	string	`json:"vlrNumber,omitempty"`
	 MscNumber	string	`json:"mscNumber,omitempty"`
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
}
