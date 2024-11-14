package models
type TraceDataResponse struct {
	 TraceData	*TraceData	`json:"traceData,omitempty"`
	 SharedTraceDataId	string	`json:"sharedTraceDataId,omitempty"`
}
