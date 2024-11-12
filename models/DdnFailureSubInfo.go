package models

type DdnFailureSubInfo struct {
	NotifyCorrelationId      string                 `json:"notifyCorrelationId"`
	DddTrafficDescriptorList []DddTrafficDescriptor `json:"dddTrafficDescriptorList,omitempty"`
}
