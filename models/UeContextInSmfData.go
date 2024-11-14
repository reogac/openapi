/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextInSmfData struct {
	PduSessions   map[string]PduSession `json:"pduSessions,omitempty"`
	PgwInfo       []PgwInfo             `json:"pgwInfo,omitempty"`
	EmergencyInfo *EmergencyInfo        `json:"emergencyInfo,omitempty"`
}
