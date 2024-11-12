package models

type NfStatus struct {
	StatusRegistered     *int `json:"statusRegistered,omitempty"`
	StatusUnregistered   *int `json:"statusUnregistered,omitempty"`
	StatusUndiscoverable *int `json:"statusUndiscoverable,omitempty"`
}
