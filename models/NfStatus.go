package models
type NfStatus struct {
	 StatusUnregistered	*int	`json:"statusUnregistered,omitempty"`
	 StatusUndiscoverable	*int	`json:"statusUndiscoverable,omitempty"`
	 StatusRegistered	*int	`json:"statusRegistered,omitempty"`
}
