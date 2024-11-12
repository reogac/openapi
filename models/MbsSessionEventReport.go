package models

type MbsSessionEventReport struct {
	EventType          MbsSessionEventType     `json:"eventType"`
	TimeStamp          string                  `json:"timeStamp,omitempty"`
	IngressTunAddrInfo *IngressTunAddrInfo     `json:"ingressTunAddrInfo,omitempty"`
	BroadcastDelStatus BroadcastDeliveryStatus `json:"broadcastDelStatus,omitempty"`
}
