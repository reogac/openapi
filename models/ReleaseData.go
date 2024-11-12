package models

type ReleaseData struct {
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	SecondaryRatUsageInfo   []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	N4Info                  *N4Information            `json:"n4Info,omitempty"`
	N4InfoExt1              *N4Information            `json:"n4InfoExt1,omitempty"`
	Cause                   string                    `json:"cause,omitempty"`
	NgApCause               *NgApCause                `json:"ngApCause,omitempty"`
	FiveGMmCauseValue       *int                      `json:"5gMmCauseValue,omitempty"`
	UeLocation              *UserLocation             `json:"ueLocation,omitempty"`
	AddUeLocation           *UserLocation             `json:"addUeLocation,omitempty"`
	SecondaryRatUsageReport []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	N4InfoExt2              *N4Information            `json:"n4InfoExt2,omitempty"`
}
