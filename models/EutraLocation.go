package models

type EutraLocation struct {
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	Tai                      Tai              `json:"tai"`
	Ecgi                     Ecgi             `json:"ecgi"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
}
