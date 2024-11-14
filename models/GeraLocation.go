package models
type GeraLocation struct {
	 Sai	*ServiceAreaId	`json:"sai,omitempty"`
	 Rai	*RoutingAreaId	`json:"rai,omitempty"`
	 VlrNumber	string	`json:"vlrNumber,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
	 LocationNumber	string	`json:"locationNumber,omitempty"`
	 Cgi	*CellGlobalId	`json:"cgi,omitempty"`
	 Lai	*LocationAreaId	`json:"lai,omitempty"`
	 MscNumber	string	`json:"mscNumber,omitempty"`
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
}
