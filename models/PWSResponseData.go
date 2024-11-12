package models

type PWSResponseData struct {
	UnknownTaiList    []Tai `json:"unknownTaiList,omitempty"`
	NgapMessageType   int   `json:"ngapMessageType"`
	SerialNumber      int   `json:"serialNumber"`
	MessageIdentifier int   `json:"messageIdentifier"`
}
