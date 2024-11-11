package models

type SendMoDataReqData struct {
	UeLocation       *UserLocation     `json:"ueLocation,omitempty"`
	MoData           RefToBinaryData   `json:"moData"`
	MoExpDataCounter *MoExpDataCounter `json:"moExpDataCounter,omitempty"`
}
