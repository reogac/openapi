package models

type EutraLocation struct {
	Tai                      Tai              `json:"tai"`
	Ecgi                     Ecgi             `json:"ecgi"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
}
