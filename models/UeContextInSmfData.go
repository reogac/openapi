package models

type UeContextInSmfData struct {
	EmergencyInfo *EmergencyInfo        `json:"emergencyInfo,omitempty"`
	PduSessions   map[string]PduSession `json:"pduSessions,omitempty"`
	PgwInfo       []PgwInfo             `json:"pgwInfo,omitempty"`
}
