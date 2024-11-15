/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleaseData struct {
	SecondaryRatUsageInfo   []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	N4Info                  *N4Information            `json:"n4Info,omitempty"`
	N4InfoExt2              *N4Information            `json:"n4InfoExt2,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	SecondaryRatUsageReport []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	FiveGMmCauseValue       *int                      `json:"5gMmCauseValue,omitempty"`
	UeLocation              *UserLocation             `json:"ueLocation,omitempty"`
	AddUeLocation           *UserLocation             `json:"addUeLocation,omitempty"`
	N4InfoExt1              *N4Information            `json:"n4InfoExt1,omitempty"`
	Cause                   Cause                     `json:"cause,omitempty"`
	NgApCause               *NgApCause                `json:"ngApCause,omitempty"`
}
