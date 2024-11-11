package models

type QosFlowProfile struct {
	NonDynamic5Qi         *NonDynamic5Qi         `json:"nonDynamic5Qi,omitempty"`
	Arp                   *Arp                   `json:"arp,omitempty"`
	GbrQosFlowInfo        *GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
	Rqa                   string                 `json:"rqa,omitempty"`
	QosMonitoringReq      string                 `json:"qosMonitoringReq,omitempty"`
	FiveQi                int                    `json:"5qi"`
	AdditionalQosFlowInfo string                 `json:"additionalQosFlowInfo,omitempty"`
	QosRepPeriod          *int                   `json:"qosRepPeriod,omitempty"`
	Dynamic5Qi            *Dynamic5Qi            `json:"dynamic5Qi,omitempty"`
}
