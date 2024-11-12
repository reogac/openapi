package models

type UeInAreaFilter struct {
	UeType          string `json:"ueType,omitempty"`
	AerialSrvDnnInd *bool  `json:"aerialSrvDnnInd,omitempty"`
}
