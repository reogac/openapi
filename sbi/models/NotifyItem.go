/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:08 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NotifyItem struct {
	ResourceId string       `json:"resourceId"`
	Changes    []ChangeItem `json:"changes"`
}