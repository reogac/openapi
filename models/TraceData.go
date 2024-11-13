package models

type TraceData struct {
	TraceDepth                TraceDepth `json:"traceDepth"`
	NeTypeList                string     `json:"neTypeList"`
	EventList                 string     `json:"eventList"`
	CollectionEntityIpv4Addr  string     `json:"collectionEntityIpv4Addr,omitempty"`
	CollectionEntityIpv6Addr  string     `json:"collectionEntityIpv6Addr,omitempty"`
	TraceReportingConsumerUri string     `json:"traceReportingConsumerUri,omitempty"`
	InterfaceList             string     `json:"interfaceList,omitempty"`
	TraceRef                  string     `json:"traceRef"`
}
