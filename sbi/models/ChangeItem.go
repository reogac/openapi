/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChangeItem struct {
	From string     `json:"from,omitempty"`
	Op   ChangeType `json:"op"`
	Path string     `json:"path"`
}
