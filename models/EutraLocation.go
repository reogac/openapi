package models

type EutraLocation struct {
	Ecgi                     Ecgi             `json:"ecgi"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
	Tai                      Tai              `json:"tai"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
}
