package models

type UeInAreaFilter struct {
	AerialSrvDnnInd *bool  `json:"aerialSrvDnnInd,omitempty"`
	UeType          string `json:"ueType,omitempty"`
}
