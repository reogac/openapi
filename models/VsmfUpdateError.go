package models
type VsmfUpdateError struct {
	 N4InfoExt3	*N4Information	`json:"n4InfoExt3,omitempty"`
	 Error	ExtProblemDetails	`json:"error"`
	 Pti	*int	`json:"pti,omitempty"`
	 N1smCause	string	`json:"n1smCause,omitempty"`
	 UnknownN1SmInfo	*RefToBinaryData	`json:"unknownN1SmInfo,omitempty"`
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
	 N1SmInfoFromUe	*RefToBinaryData	`json:"n1SmInfoFromUe,omitempty"`
	 FailedToAssignEbiList	[]Arp	`json:"failedToAssignEbiList,omitempty"`
	 FiveGMmCauseValue	*int	`json:"5gMmCauseValue,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
}
