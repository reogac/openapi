/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type V2xContext struct {
	LteUeSidelinkAmbr  string      `json:"lteUeSidelinkAmbr,omitempty"`
	Pc5QoSPara         *Pc5QoSPara `json:"pc5QoSPara,omitempty"`
	NrV2xServicesAuth  *NrV2xAuth  `json:"nrV2xServicesAuth,omitempty"`
	LteV2xServicesAuth *LteV2xAuth `json:"lteV2xServicesAuth,omitempty"`
	NrUeSidelinkAmbr   string      `json:"nrUeSidelinkAmbr,omitempty"`
}
