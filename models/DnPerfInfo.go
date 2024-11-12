package models

type DnPerfInfo struct {
	AppId      string   `json:"appId,omitempty"`
	Dnn        string   `json:"dnn,omitempty"`
	Snssai     *Snssai  `json:"snssai,omitempty"`
	DnPerf     []DnPerf `json:"dnPerf"`
	Confidence *int     `json:"confidence,omitempty"`
}
