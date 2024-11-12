package models

type UeContextInSmfData struct {
	PgwInfo       []PgwInfo             `json:"pgwInfo,omitempty"`
	EmergencyInfo *EmergencyInfo        `json:"emergencyInfo,omitempty"`
	PduSessions   map[string]PduSession `json:"pduSessions,omitempty"`
}
