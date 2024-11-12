package models

type PWSResponseData struct {
	SerialNumber      int   `json:"serialNumber"`
	MessageIdentifier int   `json:"messageIdentifier"`
	UnknownTaiList    []Tai `json:"unknownTaiList,omitempty"`
	NgapMessageType   int   `json:"ngapMessageType"`
}
