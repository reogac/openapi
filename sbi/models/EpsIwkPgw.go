/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EpsIwkPgw struct {
	PlmnId        *PlmnId `json:"plmnId,omitempty"`
	PgwFqdn       string  `json:"pgwFqdn"`
	SmfInstanceId string  `json:"smfInstanceId"`
}
