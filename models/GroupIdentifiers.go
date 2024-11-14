package models
type GroupIdentifiers struct {
	 IntGroupId	string	`json:"intGroupId,omitempty"`
	 UeIdList	[]UeId	`json:"ueIdList,omitempty"`
	 ExtGroupId	string	`json:"extGroupId,omitempty"`
}
