package models

type ReportingFrequency string

// Define constant values for ReportingFrequency
const (
	REPORTINGFREQUENCY_EVENT_TRIGGERED ReportingFrequency = "EVENT_TRIGGERED"
	REPORTINGFREQUENCY_PERIODIC        ReportingFrequency = "PERIODIC"
)
