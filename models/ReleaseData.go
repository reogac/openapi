package models

type ReleaseData struct {
	NgApCause               *NgApCause                `json:"ngApCause,omitempty"`
	FiveGMmCauseValue       *int                      `json:"5gMmCauseValue,omitempty"`
	UeLocation              *UserLocation             `json:"ueLocation,omitempty"`
	SecondaryRatUsageReport []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	N4Info                  *N4Information            `json:"n4Info,omitempty"`
	Cause                   Cause                     `json:"cause,omitempty"`
	AddUeLocation           *UserLocation             `json:"addUeLocation,omitempty"`
	SecondaryRatUsageInfo   []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	N4InfoExt1              *N4Information            `json:"n4InfoExt1,omitempty"`
	N4InfoExt2              *N4Information            `json:"n4InfoExt2,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
}
