package models

type TraceDepth string

// Define constant values for TraceDepth
const (
	TRACEDEPTH_MINIMUM                     TraceDepth = "MINIMUM"
	TRACEDEPTH_MEDIUM                      TraceDepth = "MEDIUM"
	TRACEDEPTH_MAXIMUM                     TraceDepth = "MAXIMUM"
	TRACEDEPTH_MINIMUM_WO_VENDOR_EXTENSION TraceDepth = "MINIMUM_WO_VENDOR_EXTENSION"
	TRACEDEPTH_MEDIUM_WO_VENDOR_EXTENSION  TraceDepth = "MEDIUM_WO_VENDOR_EXTENSION"
	TRACEDEPTH_MAXIMUM_WO_VENDOR_EXTENSION TraceDepth = "MAXIMUM_WO_VENDOR_EXTENSION"
)