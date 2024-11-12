package models

type EutraLocation struct {
	Ecgi                     Ecgi             `json:"ecgi"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	Tai                      Tai              `json:"tai"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
}
