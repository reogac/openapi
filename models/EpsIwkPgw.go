package models

type EpsIwkPgw struct {
	PlmnId        *PlmnId `json:"plmnId,omitempty"`
	PgwFqdn       string  `json:"pgwFqdn"`
	SmfInstanceId string  `json:"smfInstanceId"`
}
