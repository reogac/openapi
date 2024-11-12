package models
type MbsSessionEventReportList struct {
	 EventReportList	[]MbsSessionEventReport	`json:"eventReportList"`
	 NotifyCorrelationId	string	`json:"notifyCorrelationId,omitempty"`
}
