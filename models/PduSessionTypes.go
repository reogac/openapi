package models

type PduSessionTypes struct {
	AllowedSessionTypes []string       `json:"allowedSessionTypes,omitempty"`
	DefaultSessionType  PduSessionType `json:"defaultSessionType,omitempty"`
}
