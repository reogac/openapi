package models

type QosFlowProfile struct {
	FiveQi                int                    `json:"5qi"`
	NonDynamic5Qi         *NonDynamic5Qi         `json:"nonDynamic5Qi,omitempty"`
	GbrQosFlowInfo        *GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
	Dynamic5Qi            *Dynamic5Qi            `json:"dynamic5Qi,omitempty"`
	Arp                   *Arp                   `json:"arp,omitempty"`
	Rqa                   string                 `json:"rqa,omitempty"`
	AdditionalQosFlowInfo string                 `json:"additionalQosFlowInfo,omitempty"`
	QosMonitoringReq      string                 `json:"qosMonitoringReq,omitempty"`
	QosRepPeriod          *int                   `json:"qosRepPeriod,omitempty"`
}