type GeraLocation struct {
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
	 LocationNumber	string	`json:"locationNumber,omitempty"`
	 Rai	*RoutingAreaId	`json:"rai,omitempty"`
	 MscNumber	string	`json:"mscNumber,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
	 Cgi	*CellGlobalId	`json:"cgi,omitempty"`
	 Sai	*ServiceAreaId	`json:"sai,omitempty"`
	 Lai	*LocationAreaId	`json:"lai,omitempty"`
	 VlrNumber	string	`json:"vlrNumber,omitempty"`
}
