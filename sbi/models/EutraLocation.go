/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EutraLocation struct {
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	Tai                      Tai              `json:"tai"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	Ecgi                     Ecgi             `json:"ecgi"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
}
