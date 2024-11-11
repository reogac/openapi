package models

type InterFreqTargetInfo struct {
	DlCarrierFreq int   `json:"dlCarrierFreq"`
	CellIdList    []int `json:"cellIdList,omitempty"`
}
