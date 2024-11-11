type VsmfUpdateError struct {
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
	 Pti	*int	`json:"pti,omitempty"`
	 N1smCause	string	`json:"n1smCause,omitempty"`
	 UnknownN1SmInfo	*RefToBinaryData	`json:"unknownN1SmInfo,omitempty"`
	 5gMmCauseValue	*int	`json:"5gMmCauseValue,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 Error	ProblemDetails	`json:"error"`
	 N1SmInfoFromUe	*RefToBinaryData	`json:"n1SmInfoFromUe,omitempty"`
	 FailedToAssignEbiList	[]Arp	`json:"failedToAssignEbiList,omitempty"`
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
}
