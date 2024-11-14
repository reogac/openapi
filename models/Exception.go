/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Exception struct {
	ExcepTrend ExceptionTrend `json:"excepTrend,omitempty"`
	ExcepId    ExceptionId    `json:"excepId"`
	ExcepLevel *int           `json:"excepLevel,omitempty"`
}
