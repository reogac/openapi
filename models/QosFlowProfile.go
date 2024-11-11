type QosFlowProfile struct {
	 GbrQosFlowInfo	*GbrQosFlowInformation	`json:"gbrQosFlowInfo,omitempty"`
	 Rqa	string	`json:"rqa,omitempty"`
	 5qi	int	`json:"5qi"`
	 NonDynamic5Qi	*NonDynamic5Qi	`json:"nonDynamic5Qi,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
	 QosRepPeriod	*int	`json:"qosRepPeriod,omitempty"`
	 Dynamic5Qi	*Dynamic5Qi	`json:"dynamic5Qi,omitempty"`
	 AdditionalQosFlowInfo	string	`json:"additionalQosFlowInfo,omitempty"`
	 QosMonitoringReq	string	`json:"qosMonitoringReq,omitempty"`
}
