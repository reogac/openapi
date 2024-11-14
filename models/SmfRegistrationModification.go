/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistrationModification struct {
	SmfInstanceId string `json:"smfInstanceId"`
	SmfSetId      string `json:"smfSetId,omitempty"`
	PgwFqdn       string `json:"pgwFqdn,omitempty"`
}
