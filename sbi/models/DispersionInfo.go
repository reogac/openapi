/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionInfo struct {
	TsStart        string                 `json:"tsStart"`
	TsDuration     int                    `json:"tsDuration"`
	DisperCollects []DispersionCollection `json:"disperCollects"`
	DisperType     DispersionType         `json:"disperType"`
}
