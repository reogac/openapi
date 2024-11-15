/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:02 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TraceData struct {
	CollectionEntityIpv6Addr  string     `json:"collectionEntityIpv6Addr,omitempty"`
	TraceReportingConsumerUri string     `json:"traceReportingConsumerUri,omitempty"`
	InterfaceList             string     `json:"interfaceList,omitempty"`
	TraceRef                  string     `json:"traceRef"`
	TraceDepth                TraceDepth `json:"traceDepth"`
	NeTypeList                string     `json:"neTypeList"`
	EventList                 string     `json:"eventList"`
	CollectionEntityIpv4Addr  string     `json:"collectionEntityIpv4Addr,omitempty"`
}
