package models

type SdmSubsModification struct {
	MonitoredResourceUris []string `json:"monitoredResourceUris,omitempty"`
	Expires               string   `json:"expires,omitempty"`
}
