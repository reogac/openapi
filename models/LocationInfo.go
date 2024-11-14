/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LocationInfo struct {
	Loc        UserLocation `json:"loc"`
	Ratio      *int         `json:"ratio,omitempty"`
	Confidence *int         `json:"confidence,omitempty"`
}
