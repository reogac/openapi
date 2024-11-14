package models
type AppDetectionInfo struct {
	 InstanceId	string	`json:"instanceId,omitempty"`
	 SdfDescriptions	[]FlowInformation	`json:"sdfDescriptions,omitempty"`
	 AppId	string	`json:"appId"`
}
