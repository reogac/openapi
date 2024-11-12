package models

type EutraLocation struct {
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	Tai                      Tai              `json:"tai"`
	Ecgi                     Ecgi             `json:"ecgi"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
}
