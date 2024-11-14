package models
type GeraLocation struct {
	 Sai	*ServiceAreaId	`json:"sai,omitempty"`
	 VlrNumber	string	`json:"vlrNumber,omitempty"`
	 MscNumber	string	`json:"mscNumber,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
	 LocationNumber	string	`json:"locationNumber,omitempty"`
	 Lai	*LocationAreaId	`json:"lai,omitempty"`
	 Rai	*RoutingAreaId	`json:"rai,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
	 Cgi	*CellGlobalId	`json:"cgi,omitempty"`
}
