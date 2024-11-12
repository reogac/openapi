package models

type SmContextReleaseData struct {
	IsmfReleaseOnly   *bool            `json:"ismfReleaseOnly,omitempty"`
	Cause             string           `json:"cause,omitempty"`
	FiveGMmCauseValue *int             `json:"5gMmCauseValue,omitempty"`
	UeLocation        *UserLocation    `json:"ueLocation,omitempty"`
	UeTimeZone        string           `json:"ueTimeZone,omitempty"`
	VsmfReleaseOnly   *bool            `json:"vsmfReleaseOnly,omitempty"`
	N2SmInfo          *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType      string           `json:"n2SmInfoType,omitempty"`
	NgApCause         *NgApCause       `json:"ngApCause,omitempty"`
	AddUeLocation     *UserLocation    `json:"addUeLocation,omitempty"`
}
