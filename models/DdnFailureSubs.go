package models

type DdnFailureSubs struct {
	DdnFailureSubsInd      *bool               `json:"ddnFailureSubsInd,omitempty"`
	DdnFailureSubsInfoList []DdnFailureSubInfo `json:"ddnFailureSubsInfoList,omitempty"`
}
