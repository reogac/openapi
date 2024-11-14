package models
type AppDetectionInfo struct {
	 AppId	string	`json:"appId"`
	 InstanceId	string	`json:"instanceId,omitempty"`
	 SdfDescriptions	[]FlowInformation	`json:"sdfDescriptions,omitempty"`
}
