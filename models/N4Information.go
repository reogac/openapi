package models

type N4Information struct {
	N4MessageType    string           `json:"n4MessageType"`
	N4MessagePayload RefToBinaryData  `json:"n4MessagePayload"`
	N4DnaiInfo       *DnaiInformation `json:"n4DnaiInfo,omitempty"`
}
