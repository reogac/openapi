package models

type EutraLocation struct {
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	Ecgi                     Ecgi             `json:"ecgi"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
	Tai                      Tai              `json:"tai"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
}
