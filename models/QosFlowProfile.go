package models

type QosFlowProfile struct {
	Dynamic5Qi            *Dynamic5Qi            `json:"dynamic5Qi,omitempty"`
	Arp                   *Arp                   `json:"arp,omitempty"`
	GbrQosFlowInfo        *GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
	QosMonitoringReq      string                 `json:"qosMonitoringReq,omitempty"`
	NonDynamic5Qi         *NonDynamic5Qi         `json:"nonDynamic5Qi,omitempty"`
	Rqa                   string                 `json:"rqa,omitempty"`
	AdditionalQosFlowInfo string                 `json:"additionalQosFlowInfo,omitempty"`
	QosRepPeriod          *int                   `json:"qosRepPeriod,omitempty"`
	Fiveqi                int                    `json:"5qi"`
}
