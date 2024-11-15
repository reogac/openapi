/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistrationModification struct {
	SmfInstanceId string `json:"smfInstanceId"`
	SmfSetId      string `json:"smfSetId,omitempty"`
	PgwFqdn       string `json:"pgwFqdn,omitempty"`
}
