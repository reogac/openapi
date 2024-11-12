package models

type GroupIdentifiers struct {
	UeIdList   []UeId `json:"ueIdList,omitempty"`
	ExtGroupId string `json:"extGroupId,omitempty"`
	IntGroupId string `json:"intGroupId,omitempty"`
}
