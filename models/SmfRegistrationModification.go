package models

type SmfRegistrationModification struct {
	SmfSetId      string `json:"smfSetId,omitempty"`
	PgwFqdn       string `json:"pgwFqdn,omitempty"`
	SmfInstanceId string `json:"smfInstanceId"`
}
