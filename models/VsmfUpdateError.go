package models

type VsmfUpdateError struct {
	Pti                   *int             `json:"pti,omitempty"`
	UnknownN1SmInfo       *RefToBinaryData `json:"unknownN1SmInfo,omitempty"`
	N4Info                *N4Information   `json:"n4Info,omitempty"`
	N4InfoExt1            *N4Information   `json:"n4InfoExt1,omitempty"`
	Error                 ProblemDetails   `json:"error"`
	N1smCause             string           `json:"n1smCause,omitempty"`
	N1SmInfoFromUe        *RefToBinaryData `json:"n1SmInfoFromUe,omitempty"`
	FailedToAssignEbiList []Arp            `json:"failedToAssignEbiList,omitempty"`
	NgApCause             *NgApCause       `json:"ngApCause,omitempty"`
	FivegMmCauseValue     *int             `json:"5gMmCauseValue,omitempty"`
	RecoveryTime          string           `json:"recoveryTime,omitempty"`
	N4InfoExt2            *N4Information   `json:"n4InfoExt2,omitempty"`
}
