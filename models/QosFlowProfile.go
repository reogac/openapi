package models

type QosFlowProfile struct {
	NonDynamic5Qi         *NonDynamic5Qi         `json:"nonDynamic5Qi,omitempty"`
	Dynamic5Qi            *Dynamic5Qi            `json:"dynamic5Qi,omitempty"`
	Arp                   *Arp                   `json:"arp,omitempty"`
	Rqa                   ReflectiveQoSAttribute `json:"rqa,omitempty"`
	QosMonitoringReq      QosMonitoringReq       `json:"qosMonitoringReq,omitempty"`
	QosRepPeriod          *int                   `json:"qosRepPeriod,omitempty"`
	FiveQi                int                    `json:"5qi"`
	GbrQosFlowInfo        *GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
	AdditionalQosFlowInfo AdditionalQosFlowInfo  `json:"additionalQosFlowInfo,omitempty"`
}
