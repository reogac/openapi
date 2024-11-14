package models
type MbsSessionEventReport struct {
	 IngressTunAddrInfo	*IngressTunAddrInfo	`json:"ingressTunAddrInfo,omitempty"`
	 BroadcastDelStatus	BroadcastDeliveryStatus	`json:"broadcastDelStatus,omitempty"`
	 EventType	MbsSessionEventType	`json:"eventType"`
	 TimeStamp	string	`json:"timeStamp,omitempty"`
}
