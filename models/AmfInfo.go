package models

type AmfInfo struct {
	AccessType    AccessType `json:"accessType,omitempty"`
	AmfInstanceId string     `json:"amfInstanceId"`
	Guami         Guami      `json:"guami"`
}
