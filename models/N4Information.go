package models

type N4Information struct {
	UlClBpId         string           `json:"ulClBpId,omitempty"`
	N9UlPdrIdList    []int            `json:"n9UlPdrIdList,omitempty"`
	N4MessageType    string           `json:"n4MessageType"`
	N4MessagePayload RefToBinaryData  `json:"n4MessagePayload"`
	N4DnaiInfo       *DnaiInformation `json:"n4DnaiInfo,omitempty"`
	PsaUpfId         string           `json:"psaUpfId,omitempty"`
}
