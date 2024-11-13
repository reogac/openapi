package models

type AmfInfo struct {
	Guami         Guami      `json:"guami"`
	AccessType    AccessType `json:"accessType,omitempty"`
	AmfInstanceId string     `json:"amfInstanceId"`
}
