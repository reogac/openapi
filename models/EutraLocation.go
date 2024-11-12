package models

type EutraLocation struct {
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
	Tai                      Tai              `json:"tai"`
	Ecgi                     Ecgi             `json:"ecgi"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
}
