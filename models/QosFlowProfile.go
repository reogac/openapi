package models

type QosFlowProfile struct {
	GbrQosFlowInfo        *GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
	QosMonitoringReq      QosMonitoringReq       `json:"qosMonitoringReq,omitempty"`
	QosRepPeriod          *int                   `json:"qosRepPeriod,omitempty"`
	FiveQi                int                    `json:"5qi"`
	NonDynamic5Qi         *NonDynamic5Qi         `json:"nonDynamic5Qi,omitempty"`
	Rqa                   ReflectiveQoSAttribute `json:"rqa,omitempty"`
	AdditionalQosFlowInfo AdditionalQosFlowInfo  `json:"additionalQosFlowInfo,omitempty"`
	Dynamic5Qi            *Dynamic5Qi            `json:"dynamic5Qi,omitempty"`
	Arp                   *Arp                   `json:"arp,omitempty"`
}
