package models

type ReleaseData struct {
	N4InfoExt2              *N4Information            `json:"n4InfoExt2,omitempty"`
	Cause                   Cause                     `json:"cause,omitempty"`
	UeLocation              *UserLocation             `json:"ueLocation,omitempty"`
	SecondaryRatUsageInfo   []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	AddUeLocation           *UserLocation             `json:"addUeLocation,omitempty"`
	SecondaryRatUsageReport []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	N4Info                  *N4Information            `json:"n4Info,omitempty"`
	N4InfoExt1              *N4Information            `json:"n4InfoExt1,omitempty"`
	NgApCause               *NgApCause                `json:"ngApCause,omitempty"`
	FiveGMmCauseValue       *int                      `json:"5gMmCauseValue,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
}
