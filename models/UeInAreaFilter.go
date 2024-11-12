package models

type UeInAreaFilter struct {
	UeType          UeType `json:"ueType,omitempty"`
	AerialSrvDnnInd *bool  `json:"aerialSrvDnnInd,omitempty"`
}
