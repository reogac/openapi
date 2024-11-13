package models

type TraceDataResponse struct {
	SharedTraceDataId string     `json:"sharedTraceDataId,omitempty"`
	TraceData         *TraceData `json:"traceData,omitempty"`
}
