package models

type NwdafFailureCode string

// Define constant values for NwdafFailureCode
const (
	NWDAFFAILURECODE_UNAVAILABLE_DATA                     NwdafFailureCode = "UNAVAILABLE_DATA"
	NWDAFFAILURECODE_BOTH_STAT_PRED_NOT_ALLOWED           NwdafFailureCode = "BOTH_STAT_PRED_NOT_ALLOWED"
	NWDAFFAILURECODE_UNSATISFIED_REQUESTED_ANALYTICS_TIME NwdafFailureCode = "UNSATISFIED_REQUESTED_ANALYTICS_TIME"
	NWDAFFAILURECODE_OTHER                                NwdafFailureCode = "OTHER"
)