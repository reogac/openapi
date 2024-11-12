package models
type MbsSessionEventReport struct {
	 EventType	string	`json:"eventType"`
	 TimeStamp	string	`json:"timeStamp,omitempty"`
	 IngressTunAddrInfo	*IngressTunAddrInfo	`json:"ingressTunAddrInfo,omitempty"`
	 BroadcastDelStatus	string	`json:"broadcastDelStatus,omitempty"`
}
