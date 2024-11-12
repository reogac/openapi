package models
type EutraLocation struct {
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 IgnoreTai	*bool	`json:"ignoreTai,omitempty"`
	 Ecgi	Ecgi	`json:"ecgi"`
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
	 GlobalNgenbId	*GlobalRanNodeId	`json:"globalNgenbId,omitempty"`
	 GlobalENbId	*GlobalRanNodeId	`json:"globalENbId,omitempty"`
	 Tai	Tai	`json:"tai"`
	 IgnoreEcgi	*bool	`json:"ignoreEcgi,omitempty"`
}
