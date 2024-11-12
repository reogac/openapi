package models
type EpsIwkPgw struct {
	 SmfInstanceId	string	`json:"smfInstanceId"`
	 PlmnId	*PlmnId	`json:"plmnId,omitempty"`
	 PgwFqdn	string	`json:"pgwFqdn"`
}
