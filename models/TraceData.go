package models

type TraceData struct {
	EventList                string `json:"eventList"`
	CollectionEntityIpv4Addr string `json:"collectionEntityIpv4Addr,omitempty"`
	CollectionEntityIpv6Addr string `json:"collectionEntityIpv6Addr,omitempty"`
	InterfaceList            string `json:"interfaceList,omitempty"`
	TraceRef                 string `json:"traceRef"`
	TraceDepth               string `json:"traceDepth"`
	NeTypeList               string `json:"neTypeList"`
}
