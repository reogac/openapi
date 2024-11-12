package models

type QosResourceType string

// Define constant values for QosResourceType
const (
	QOSRESOURCETYPE_NON_GBR          QosResourceType = "NON_GBR"
	QOSRESOURCETYPE_NON_CRITICAL_GBR QosResourceType = "NON_CRITICAL_GBR"
	QOSRESOURCETYPE_CRITICAL_GBR     QosResourceType = "CRITICAL_GBR"
)
