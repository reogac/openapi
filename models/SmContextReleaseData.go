package models

type SmContextReleaseData struct {
	UeLocation        *UserLocation    `json:"ueLocation,omitempty"`
	AddUeLocation     *UserLocation    `json:"addUeLocation,omitempty"`
	VsmfReleaseOnly   *bool            `json:"vsmfReleaseOnly,omitempty"`
	N2SmInfoType      string           `json:"n2SmInfoType,omitempty"`
	Cause             string           `json:"cause,omitempty"`
	NgApCause         *NgApCause       `json:"ngApCause,omitempty"`
	FiveGMmCauseValue *int             `json:"5gMmCauseValue,omitempty"`
	UeTimeZone        string           `json:"ueTimeZone,omitempty"`
	N2SmInfo          *RefToBinaryData `json:"n2SmInfo,omitempty"`
	IsmfReleaseOnly   *bool            `json:"ismfReleaseOnly,omitempty"`
}
