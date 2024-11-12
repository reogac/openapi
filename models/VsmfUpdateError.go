package models

type VsmfUpdateError struct {
	N4InfoExt3            *N4Information    `json:"n4InfoExt3,omitempty"`
	N1SmInfoFromUe        *RefToBinaryData  `json:"n1SmInfoFromUe,omitempty"`
	UnknownN1SmInfo       *RefToBinaryData  `json:"unknownN1SmInfo,omitempty"`
	NgApCause             *NgApCause        `json:"ngApCause,omitempty"`
	N4Info                *N4Information    `json:"n4Info,omitempty"`
	N4InfoExt2            *N4Information    `json:"n4InfoExt2,omitempty"`
	RecoveryTime          string            `json:"recoveryTime,omitempty"`
	N4InfoExt1            *N4Information    `json:"n4InfoExt1,omitempty"`
	Error                 ExtProblemDetails `json:"error"`
	Pti                   *int              `json:"pti,omitempty"`
	N1smCause             string            `json:"n1smCause,omitempty"`
	FailedToAssignEbiList []Arp             `json:"failedToAssignEbiList,omitempty"`
	FiveGMmCauseValue     *int              `json:"5gMmCauseValue,omitempty"`
}
