package models

type PgwInfo struct {
	PgwFqdn          string     `json:"pgwFqdn"`
	PgwIpAddr        *IpAddress `json:"pgwIpAddr,omitempty"`
	PlmnId           *PlmnId    `json:"plmnId,omitempty"`
	EpdgInd          *bool      `json:"epdgInd,omitempty"`
	PcfId            string     `json:"pcfId,omitempty"`
	RegistrationTime string     `json:"registrationTime,omitempty"`
	Dnn              string     `json:"dnn"`
}
