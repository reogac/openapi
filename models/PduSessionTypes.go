package models
type PduSessionTypes struct {
	 DefaultSessionType	PduSessionType	`json:"defaultSessionType,omitempty"`
	 AllowedSessionTypes	[]string	`json:"allowedSessionTypes,omitempty"`
}
