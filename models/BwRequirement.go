/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BwRequirement struct {
	AppId   string `json:"appId"`
	MarBwDl string `json:"marBwDl,omitempty"`
	MarBwUl string `json:"marBwUl,omitempty"`
	MirBwDl string `json:"mirBwDl,omitempty"`
	MirBwUl string `json:"mirBwUl,omitempty"`
}
