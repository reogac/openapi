/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type V2xSubscriptionData struct {
	NrV2xServicesAuth  *NrV2xAuth  `json:"nrV2xServicesAuth,omitempty"`
	LteV2xServicesAuth *LteV2xAuth `json:"lteV2xServicesAuth,omitempty"`
	NrUePc5Ambr        string      `json:"nrUePc5Ambr,omitempty"`
	LtePc5Ambr         string      `json:"ltePc5Ambr,omitempty"`
}
