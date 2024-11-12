package models

type SmcceInfo struct {
	Snssai      *Snssai     `json:"snssai,omitempty"`
	SmcceUeList SmcceUeList `json:"smcceUeList"`
	Dnn         string      `json:"dnn,omitempty"`
}
