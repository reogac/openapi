package models

type SendMoDataReqData struct {
	MoData           RefToBinaryData   `json:"moData"`
	MoExpDataCounter *MoExpDataCounter `json:"moExpDataCounter,omitempty"`
	UeLocation       *UserLocation     `json:"ueLocation,omitempty"`
}
