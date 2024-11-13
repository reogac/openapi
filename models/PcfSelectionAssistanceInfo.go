package models

type PcfSelectionAssistanceInfo struct {
	SingleNssai Snssai `json:"singleNssai"`
	Dnn         string `json:"dnn"`
}
