package models

type DdnFailureSubInfo struct {
	DddTrafficDescriptorList []DddTrafficDescriptor `json:"dddTrafficDescriptorList,omitempty"`
	NotifyCorrelationId      string                 `json:"notifyCorrelationId"`
}
