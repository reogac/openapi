package models
type QosFlowProfile struct {
	 Rqa	ReflectiveQoSAttribute	`json:"rqa,omitempty"`
	 QosMonitoringReq	QosMonitoringReq	`json:"qosMonitoringReq,omitempty"`
	 FiveQi	int	`json:"5qi"`
	 Dynamic5Qi	*Dynamic5Qi	`json:"dynamic5Qi,omitempty"`
	 GbrQosFlowInfo	*GbrQosFlowInformation	`json:"gbrQosFlowInfo,omitempty"`
	 AdditionalQosFlowInfo	AdditionalQosFlowInfo	`json:"additionalQosFlowInfo,omitempty"`
	 QosRepPeriod	*int	`json:"qosRepPeriod,omitempty"`
	 NonDynamic5Qi	*NonDynamic5Qi	`json:"nonDynamic5Qi,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
}
