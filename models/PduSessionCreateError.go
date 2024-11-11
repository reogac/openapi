type PduSessionCreateError struct {
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
	 Error	ProblemDetails	`json:"error"`
	 N1smCause	string	`json:"n1smCause,omitempty"`
	 N1SmInfoToUe	*RefToBinaryData	`json:"n1SmInfoToUe,omitempty"`
	 BackOffTimer	*int	`json:"backOffTimer,omitempty"`
}
