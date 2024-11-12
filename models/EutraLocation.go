package models
type EutraLocation struct {
	 Tai	Tai	`json:"tai"`
	 IgnoreTai	*bool	`json:"ignoreTai,omitempty"`
	 Ecgi	Ecgi	`json:"ecgi"`
	 IgnoreEcgi	*bool	`json:"ignoreEcgi,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
	 GlobalNgenbId	*GlobalRanNodeId	`json:"globalNgenbId,omitempty"`
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GlobalENbId	*GlobalRanNodeId	`json:"globalENbId,omitempty"`
}
