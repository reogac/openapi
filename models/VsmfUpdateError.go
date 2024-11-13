package models

type VsmfUpdateError struct {
	N1smCause             string            `json:"n1smCause,omitempty"`
	UnknownN1SmInfo       *RefToBinaryData  `json:"unknownN1SmInfo,omitempty"`
	FailedToAssignEbiList []Arp             `json:"failedToAssignEbiList,omitempty"`
	NgApCause             *NgApCause        `json:"ngApCause,omitempty"`
	FiveGMmCauseValue     *int              `json:"5gMmCauseValue,omitempty"`
	RecoveryTime          string            `json:"recoveryTime,omitempty"`
	Error                 ExtProblemDetails `json:"error"`
	Pti                   *int              `json:"pti,omitempty"`
	N4InfoExt3            *N4Information    `json:"n4InfoExt3,omitempty"`
	N4Info                *N4Information    `json:"n4Info,omitempty"`
	N4InfoExt2            *N4Information    `json:"n4InfoExt2,omitempty"`
	N1SmInfoFromUe        *RefToBinaryData  `json:"n1SmInfoFromUe,omitempty"`
	N4InfoExt1            *N4Information    `json:"n4InfoExt1,omitempty"`
}
