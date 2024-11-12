package models
type NrLocation struct {
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
	 GlobalGnbId	*GlobalRanNodeId	`json:"globalGnbId,omitempty"`
	 Tai	Tai	`json:"tai"`
	 Ncgi	Ncgi	`json:"ncgi"`
	 IgnoreNcgi	*bool	`json:"ignoreNcgi,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
}
