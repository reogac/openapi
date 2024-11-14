/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionInfo struct {
	TsDuration     int                    `json:"tsDuration"`
	DisperCollects []DispersionCollection `json:"disperCollects"`
	DisperType     DispersionType         `json:"disperType"`
	TsStart        string                 `json:"tsStart"`
}
