package models

type RequestType string

// Define constant values for RequestType
const (
	REQUESTTYPE_INITIAL_REQUEST                RequestType = "INITIAL_REQUEST"
	REQUESTTYPE_EXISTING_PDU_SESSION           RequestType = "EXISTING_PDU_SESSION"
	REQUESTTYPE_INITIAL_EMERGENCY_REQUEST      RequestType = "INITIAL_EMERGENCY_REQUEST"
	REQUESTTYPE_EXISTING_EMERGENCY_PDU_SESSION RequestType = "EXISTING_EMERGENCY_PDU_SESSION"
)