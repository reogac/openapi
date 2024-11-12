package models

type InterFreqTargetInfo struct {
	CellIdList    []int `json:"cellIdList,omitempty"`
	DlCarrierFreq int   `json:"dlCarrierFreq"`
}
