package models
type CongestionType string
// Define constant values for CongestionType
const (
	 CONGESTIONTYPE_USER_PLANE CongestionType = "USER_PLANE"
	 CONGESTIONTYPE_CONTROL_PLANE CongestionType = "CONTROL_PLANE"
	 CONGESTIONTYPE_USER_AND_CONTROL_PLANE CongestionType = "USER_AND_CONTROL_PLANE"
) 
