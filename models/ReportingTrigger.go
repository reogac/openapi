package models

type ReportingTrigger string

// Define constant values for ReportingTrigger
const (
	REPORTINGTRIGGER_PERIODICAL             ReportingTrigger = "PERIODICAL"
	REPORTINGTRIGGER_EVENT_A2               ReportingTrigger = "EVENT_A2"
	REPORTINGTRIGGER_EVENT_A2_PERIODIC      ReportingTrigger = "EVENT_A2_PERIODIC"
	REPORTINGTRIGGER_ALL_RRM_EVENT_TRIGGERS ReportingTrigger = "ALL_RRM_EVENT_TRIGGERS"
)
