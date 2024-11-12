package models

type AppDetectionInfo struct {
	SdfDescriptions []FlowInformation `json:"sdfDescriptions,omitempty"`
	AppId           string            `json:"appId"`
	InstanceId      string            `json:"instanceId,omitempty"`
}
