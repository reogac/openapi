package models

type DnPerfInfo struct {
	DnPerf     []DnPerf `json:"dnPerf"`
	Confidence *int     `json:"confidence,omitempty"`
	AppId      string   `json:"appId,omitempty"`
	Dnn        string   `json:"dnn,omitempty"`
	Snssai     *Snssai  `json:"snssai,omitempty"`
}
