package models

type TransferMoDataReqData struct {
	MoExpDataCounter *MoExpDataCounter `json:"moExpDataCounter,omitempty"`
	UeLocation       *UserLocation     `json:"ueLocation,omitempty"`
	MoData           RefToBinaryData   `json:"moData"`
}
