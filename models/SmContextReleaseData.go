package models

type SmContextReleaseData struct {
	Cause             Cause            `json:"cause,omitempty"`
	FiveGMmCauseValue *int             `json:"5gMmCauseValue,omitempty"`
	UeTimeZone        string           `json:"ueTimeZone,omitempty"`
	N2SmInfo          *RefToBinaryData `json:"n2SmInfo,omitempty"`
	IsmfReleaseOnly   *bool            `json:"ismfReleaseOnly,omitempty"`
	NgApCause         *NgApCause       `json:"ngApCause,omitempty"`
	UeLocation        *UserLocation    `json:"ueLocation,omitempty"`
	AddUeLocation     *UserLocation    `json:"addUeLocation,omitempty"`
	VsmfReleaseOnly   *bool            `json:"vsmfReleaseOnly,omitempty"`
	N2SmInfoType      N2SmInfoType     `json:"n2SmInfoType,omitempty"`
}
