package models

type LadnInfo struct {
	Presence PresenceState `json:"presence,omitempty"`
	Ladn     string        `json:"ladn"`
}
