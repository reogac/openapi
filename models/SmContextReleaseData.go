package models

type SmContextReleaseData struct {
	NgApCause         *NgApCause       `json:"ngApCause,omitempty"`
	UeTimeZone        string           `json:"ueTimeZone,omitempty"`
	VsmfReleaseOnly   *bool            `json:"vsmfReleaseOnly,omitempty"`
	N2SmInfo          *RefToBinaryData `json:"n2SmInfo,omitempty"`
	Cause             Cause            `json:"cause,omitempty"`
	UeLocation        *UserLocation    `json:"ueLocation,omitempty"`
	AddUeLocation     *UserLocation    `json:"addUeLocation,omitempty"`
	N2SmInfoType      N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	IsmfReleaseOnly   *bool            `json:"ismfReleaseOnly,omitempty"`
	FiveGMmCauseValue *int             `json:"5gMmCauseValue,omitempty"`
}
