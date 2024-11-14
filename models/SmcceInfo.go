/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmcceInfo struct {
	Dnn         string      `json:"dnn,omitempty"`
	Snssai      *Snssai     `json:"snssai,omitempty"`
	SmcceUeList SmcceUeList `json:"smcceUeList"`
}
