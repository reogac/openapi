package models

type SdmSubsModification struct {
	Expires               string   `json:"expires,omitempty"`
	MonitoredResourceUris []string `json:"monitoredResourceUris,omitempty"`
}
