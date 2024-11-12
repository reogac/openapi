package models

type UeInAreaFilter struct {
	AerialSrvDnnInd *bool  `json:"aerialSrvDnnInd,omitempty"`
	UeType          UeType `json:"ueType,omitempty"`
}
