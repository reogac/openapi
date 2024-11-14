/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsQoSReq struct {
	ReqMbsArp   *Arp   `json:"reqMbsArp,omitempty"`
	FiveQi      int    `json:"5qi"`
	GuarBitRate string `json:"guarBitRate,omitempty"`
	MaxBitRate  string `json:"maxBitRate,omitempty"`
	AverWindow  *int   `json:"averWindow,omitempty"`
}
