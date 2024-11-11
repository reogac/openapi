package models
type PWSResponseData struct {
	 NgapMessageType	int	`json:"ngapMessageType"`
	 SerialNumber	int	`json:"serialNumber"`
	 MessageIdentifier	int	`json:"messageIdentifier"`
	 UnknownTaiList	[]Tai	`json:"unknownTaiList,omitempty"`
}
