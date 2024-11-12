package models

type SmContextReleaseData struct {
	UeTimeZone        string           `json:"ueTimeZone,omitempty"`
	AddUeLocation     *UserLocation    `json:"addUeLocation,omitempty"`
	N2SmInfoType      N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	Cause             Cause            `json:"cause,omitempty"`
	NgApCause         *NgApCause       `json:"ngApCause,omitempty"`
	FiveGMmCauseValue *int             `json:"5gMmCauseValue,omitempty"`
	IsmfReleaseOnly   *bool            `json:"ismfReleaseOnly,omitempty"`
	UeLocation        *UserLocation    `json:"ueLocation,omitempty"`
	VsmfReleaseOnly   *bool            `json:"vsmfReleaseOnly,omitempty"`
	N2SmInfo          *RefToBinaryData `json:"n2SmInfo,omitempty"`
}
