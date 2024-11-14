package models
type LadnInfo struct {
	 Ladn	string	`json:"ladn"`
	 Presence	PresenceState	`json:"presence,omitempty"`
}
