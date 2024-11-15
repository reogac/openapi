/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficCharacterization struct {
	FDescs        []IpEthFlowDescription `json:"fDescs,omitempty"`
	UlVol         *int64                 `json:"ulVol,omitempty"`
	UlVolVariance *float64               `json:"ulVolVariance,omitempty"`
	DlVol         *int64                 `json:"dlVol,omitempty"`
	DlVolVariance *float64               `json:"dlVolVariance,omitempty"`
	Dnn           string                 `json:"dnn,omitempty"`
	Snssai        *Snssai                `json:"snssai,omitempty"`
	AppId         string                 `json:"appId,omitempty"`
}
