package models
type NfStatus struct {
	 StatusUndiscoverable	*int	`json:"statusUndiscoverable,omitempty"`
	 StatusRegistered	*int	`json:"statusRegistered,omitempty"`
	 StatusUnregistered	*int	`json:"statusUnregistered,omitempty"`
}
