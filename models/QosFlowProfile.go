package models

type QosFlowProfile struct {
	FiveQi                int                    `json:"5qi"`
	NonDynamic5Qi         *NonDynamic5Qi         `json:"nonDynamic5Qi,omitempty"`
	Dynamic5Qi            *Dynamic5Qi            `json:"dynamic5Qi,omitempty"`
	AdditionalQosFlowInfo AdditionalQosFlowInfo  `json:"additionalQosFlowInfo,omitempty"`
	QosMonitoringReq      QosMonitoringReq       `json:"qosMonitoringReq,omitempty"`
	Arp                   *Arp                   `json:"arp,omitempty"`
	GbrQosFlowInfo        *GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
	Rqa                   ReflectiveQoSAttribute `json:"rqa,omitempty"`
	QosRepPeriod          *int                   `json:"qosRepPeriod,omitempty"`
}
