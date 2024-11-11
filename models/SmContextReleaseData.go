package models

type SmContextReleaseData struct {
	N2SmInfo          *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType      string           `json:"n2SmInfoType,omitempty"`
	IsmfReleaseOnly   *bool            `json:"ismfReleaseOnly,omitempty"`
	Cause             string           `json:"cause,omitempty"`
	NgApCause         *NgApCause       `json:"ngApCause,omitempty"`
	FivegMmCauseValue *int             `json:"5gMmCauseValue,omitempty"`
	AddUeLocation     *UserLocation    `json:"addUeLocation,omitempty"`
	UeLocation        *UserLocation    `json:"ueLocation,omitempty"`
	UeTimeZone        string           `json:"ueTimeZone,omitempty"`
	VsmfReleaseOnly   *bool            `json:"vsmfReleaseOnly,omitempty"`
}
