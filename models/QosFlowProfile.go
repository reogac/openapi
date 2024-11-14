package models
type QosFlowProfile struct {
	 QosRepPeriod	*int	`json:"qosRepPeriod,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
	 Rqa	ReflectiveQoSAttribute	`json:"rqa,omitempty"`
	 AdditionalQosFlowInfo	AdditionalQosFlowInfo	`json:"additionalQosFlowInfo,omitempty"`
	 GbrQosFlowInfo	*GbrQosFlowInformation	`json:"gbrQosFlowInfo,omitempty"`
	 QosMonitoringReq	QosMonitoringReq	`json:"qosMonitoringReq,omitempty"`
	 FiveQi	int	`json:"5qi"`
	 NonDynamic5Qi	*NonDynamic5Qi	`json:"nonDynamic5Qi,omitempty"`
	 Dynamic5Qi	*Dynamic5Qi	`json:"dynamic5Qi,omitempty"`
}
