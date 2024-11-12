package models

type PWSResponseData struct {
	MessageIdentifier int   `json:"messageIdentifier"`
	UnknownTaiList    []Tai `json:"unknownTaiList,omitempty"`
	NgapMessageType   int   `json:"ngapMessageType"`
	SerialNumber      int   `json:"serialNumber"`
}
