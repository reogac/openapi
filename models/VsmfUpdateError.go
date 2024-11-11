package models
type VsmfUpdateError struct {
	 FailedToAssignEbiList	[]Arp	`json:"failedToAssignEbiList,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 Error	ProblemDetails	`json:"error"`
	 Pti	*int	`json:"pti,omitempty"`
	 UnknownN1SmInfo	*RefToBinaryData	`json:"unknownN1SmInfo,omitempty"`
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
	 FiveGMmCauseValue	*int	`json:"5gMmCauseValue,omitempty"`
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 N1smCause	string	`json:"n1smCause,omitempty"`
	 N1SmInfoFromUe	*RefToBinaryData	`json:"n1SmInfoFromUe,omitempty"`
}
