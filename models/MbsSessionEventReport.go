/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsSessionEventReport struct {
	EventType          MbsSessionEventType     `json:"eventType"`
	TimeStamp          string                  `json:"timeStamp,omitempty"`
	IngressTunAddrInfo *IngressTunAddrInfo     `json:"ingressTunAddrInfo,omitempty"`
	BroadcastDelStatus BroadcastDeliveryStatus `json:"broadcastDelStatus,omitempty"`
}
