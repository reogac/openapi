package models

type QosMonitoringReq string

// Define constant values for QosMonitoringReq
const (
	QOSMONITORINGREQ_UL   QosMonitoringReq = "UL"
	QOSMONITORINGREQ_DL   QosMonitoringReq = "DL"
	QOSMONITORINGREQ_BOTH QosMonitoringReq = "BOTH"
	QOSMONITORINGREQ_NONE QosMonitoringReq = "NONE"
)
