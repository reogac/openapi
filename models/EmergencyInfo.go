package models

type EmergencyInfo struct {
	SmfInstanceId string     `json:"smfInstanceId,omitempty"`
	EpdgInd       *bool      `json:"epdgInd,omitempty"`
	PlmnId        *PlmnId    `json:"plmnId,omitempty"`
	PgwFqdn       string     `json:"pgwFqdn,omitempty"`
	PgwIpAddress  *IpAddress `json:"pgwIpAddress,omitempty"`
}
