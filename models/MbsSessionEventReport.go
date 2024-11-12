package models

type MbsSessionEventReport struct {
	TimeStamp          string                  `json:"timeStamp,omitempty"`
	IngressTunAddrInfo *IngressTunAddrInfo     `json:"ingressTunAddrInfo,omitempty"`
	BroadcastDelStatus BroadcastDeliveryStatus `json:"broadcastDelStatus,omitempty"`
	EventType          MbsSessionEventType     `json:"eventType"`
}
