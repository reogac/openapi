package models
type HsmfUpdateError struct {
	 Error	ProblemDetails	`json:"error"`
	 Pti	*int	`json:"pti,omitempty"`
	 N1smCause	string	`json:"n1smCause,omitempty"`
	 N1SmInfoToUe	*RefToBinaryData	`json:"n1SmInfoToUe,omitempty"`
	 BackOffTimer	*int	`json:"backOffTimer,omitempty"`
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
}