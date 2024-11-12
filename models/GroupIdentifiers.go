package models

type GroupIdentifiers struct {
	ExtGroupId string `json:"extGroupId,omitempty"`
	IntGroupId string `json:"intGroupId,omitempty"`
	UeIdList   []UeId `json:"ueIdList,omitempty"`
}
