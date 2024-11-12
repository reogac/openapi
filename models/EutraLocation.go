package models

type EutraLocation struct {
	Tai                      Tai              `json:"tai"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	Ecgi                     Ecgi             `json:"ecgi"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
}
