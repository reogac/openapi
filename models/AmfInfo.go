package models

type AmfInfo struct {
	AmfInstanceId string     `json:"amfInstanceId"`
	Guami         Guami      `json:"guami"`
	AccessType    AccessType `json:"accessType,omitempty"`
}
